# brainfuck_to_js_translator
Translator of Brainfuck to Javascript

# How to run
set GOPATH=.
go run src/main/main.go test/data/HelloWorld.bf test/data/HelloWorld.in
node tmp/translated.js

# How to test
go test -v ./...