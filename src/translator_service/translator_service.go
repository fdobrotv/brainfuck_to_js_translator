package translator_service

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"stack"
)

func bfJumps(prog []byte) (map[uint]uint, error) {
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

func Translate(prog []byte, input *bufio.Reader, resultFile string) error {
	var ioWriter io.Writer = os.Stdout

	var (
		fpos uint   = 0                  // file position
		dpos uint   = 0                  // data position
		size uint   = 30000              // size of data card
		plen uint   = uint(len(prog))    // programme length
		data []byte = make([]byte, size) // data card with `size` items
	)

	jumps, err := bfJumps(prog) // pre-computed jumps

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
		var jsBlock string
		switch prog[fpos] {
		case '+': // increment at current position
			jsBlock = "data[dpos] += 1;\n"
			data[dpos]++
		case '-': // decrement at current position
			jsBlock = "data[dpos] -= 1;\n"
			data[dpos]--
		case '>': // move to next position
			jsBlock = `if (dpos === size-1) {
							dpos = 0;
						} else {
							dpos++;
						}` + "\n"
			if dpos == size-1 {
				dpos = 0
			} else {
				dpos++
			}
		case '<': // move to previous position
			jsBlock = `if (dpos === 0) {
							dpos = size - 1;
						} else {
							dpos--;
						}` + "\n"
			if dpos == 0 {
				dpos = size - 1
			} else {
				dpos--
			}
		case '.': // output value of current position
			jsBlock = "process.stdout.write(String.fromCharCode(data[dpos]));\n"
			fmt.Fprintf(ioWriter, "%c", data[dpos])
		case ',': // read value into current position
			readByte, _ := input.ReadByte()
			jsBlock = fmt.Sprintf("if (data[dpos] = %d) {;", readByte) +
				`	process.exit(0)
										}` + "\n"
			data[dpos] = readByte
			if err != nil && err != io.EOF {
				os.Exit(0)
			}
		case '[': // if current position is false, skip to ]
			jsBlock = `if (data[dpos] === 0) {
						fpos = jumps[fpos]
					}` + "\n"
			if data[dpos] == 0 {
				fpos = jumps[fpos]
			}
		case ']': // if at current position true, return to [
			jsBlock = `if (data[dpos] !== 0) {
						fpos = jumps[fpos]
					}` + "\n"
			if data[dpos] != 0 {
				fpos = jumps[fpos]
			}
		}
		sb.WriteString(jsBlock)
		fpos++
	}

	d1 := []byte(sb.String())
	err2 := os.WriteFile(resultFile, d1, 0644)
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
