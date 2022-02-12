package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"stack"
)

func bf_jumps(prog []byte) (map[uint]uint, error) {
	var (
		stack *stack.Stack  = stack.New()
		jumps map[uint]uint = make(map[uint]uint)

		plen uint = uint(len(prog))
		fpos uint = 0
	)

	for fpos < plen {
		switch prog[fpos] {
		case '[':
			stack.Push(fpos)
		case ']':
			tget, err := stack.Pop()
			if err != nil {
				return nil, errors.New(
					"unexpected closing bracket",
				)
			}
			jumps[tget] = fpos
			jumps[fpos] = tget
		}
		fpos++
	}

	_, err := stack.Pop()
	if err == nil {
		return nil, errors.New(
			"excessive opening brackets",
		)
	}

	return jumps, nil
}

func translate(r io.Reader, i io.Reader, w io.Writer) error {
	prog, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	input := bufio.NewReader(i) // buffered reader for `,` requests

	var (
		fpos uint   = 0                  // file position
		dpos uint   = 0                  // data position
		size uint   = 30000              // size of data card
		plen uint   = uint(len(prog))    // programme length
		data []byte = make([]byte, size) // data card with `size` items
	)

	jumps, err := bf_jumps(prog) // pre-computed jumps

	if err != nil {
		return err
	}

	var sb strings.Builder
	fposJS := fmt.Sprintf("let fpos = %d;\n", fpos)
	sb.WriteString(fposJS)
	dposJS := fmt.Sprintf("let dpos = %d;\n", dpos)
	sb.WriteString(dposJS)
	var jsProgArray = toJSProgString(prog)
	sb.WriteString(jsProgArray)
	dataJS := fmt.Sprintf("const data = new Uint8Array(%d);\n", size)
	sb.WriteString(dataJS)
	sizeJS := fmt.Sprintf("const size = %d;\n", size)
	sb.WriteString(sizeJS)
	jumpsJS := fmt.Sprintf("const jumps = new Map([%s]);\n", toJSJumps(jumps))
	sb.WriteString(jumpsJS)

	for fpos < plen {
		var jsBlock = toJSBlock(prog[fpos], input)
		sb.WriteString(jsBlock)
		switch prog[fpos] {
		case '+': // increment at current position
			data[dpos]++
		case '-': // decrement at current position
			data[dpos]--
		case '>': // move to next position
			if dpos == size-1 {
				dpos = 0
			} else {
				dpos++
			}
		case '<': // move to previous position
			if dpos == 0 {
				dpos = size - 1
			} else {
				dpos--
			}
		case '.': // output value of current position
			fmt.Fprintf(w, "%c", data[dpos])
		case ',': // read value into current position
			if data[dpos], err = input.ReadByte(); err != nil {
				os.Exit(0)
			}
		case '[': // if current position is false, skip to ]
			if data[dpos] == 0 {
				fpos = jumps[fpos]
			}
		case ']': // if at current position true, return to [
			if data[dpos] != 0 {
				fpos = jumps[fpos]
			}
		}
		fpos++
	}

	d1 := []byte(sb.String())
	err2 := os.WriteFile("tmp/translated.js", d1, 0644)
	check(err2)

	return nil
}

func toJSJumps(jumps map[uint]uint) string {
	var preResult []string

	for key, element := range jumps {
		fmt.Println("Key:", key, "=>", "Element:", element)
		mapEntry := fmt.Sprintf("['%d', '%d']", key, element)
		preResult = append(preResult, mapEntry)
	}
	return strings.Join(preResult, ",")
}

func toJSProgString(prog []byte) string {
	onlyBrainFuckSymbols := filterBrainFuck(prog)
	fmt.Println(string(onlyBrainFuckSymbols))
	return "const prog = \"" + string(onlyBrainFuckSymbols) + "\";\n"
}

func filterBrainFuck(input []byte) (ret []byte) {
	for _, el := range input {
		if el == '+' ||
			el == '-' ||
			el == '<' ||
			el == '>' ||
			el == '.' ||
			el == ',' ||
			el == '[' ||
			el == ']' {
			ret = append(ret, el)
		}
	}
	return
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [file.bf]\n", os.Args[0])
		os.Exit(3)
	}

	r, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	err = translate(r, os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toJSBlock(char byte, input *bufio.Reader) string {
	var result string

	switch char {
	case '+': // increment at current position
		result = "data[dpos] += 1;\n"
	case '-': // decrement at current position
		result = "data[dpos] -= 1;\n"
	case '>': // move to next position
		result = `if (dpos === size-1) {
			dpos = 0;
		} else {
			dpos++;
		}` + "\n"
	case '<': // move to previous position
		result = `if (dpos === 0) {
			dpos = size - 1;
		} else {
			dpos--;
		}` + "\n"
	case '.': // output value of current position
		result = `process.stdout.write(String.fromCharCode(data[dpos]));` + "\n"
	case ',': // read value into current position
		readByte, err := input.ReadByte()
		if err != nil {
			os.Exit(0)
		}
		result = fmt.Sprintf("if (data[dpos] = %s) {;", readByte) +
			`	process.exit(0)
		}` + "\n"
	case '[': // if current position is false, skip to ]
		result = `if (data[dpos] === 0) {
			fpos = jumps[fpos]
		}` + "\n"
	case ']': // if at current position true, return to [
		result = `if (data[dpos] !== 0) {
			fpos = jumps[fpos]
		}` + "\n"
	}

	return result
}
