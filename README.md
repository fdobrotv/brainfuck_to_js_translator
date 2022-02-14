# brainfuck_to_js_translator
Translator of Brainfuck to Javascript

# How to run
set GOPATH=.
go run src/main/main.go test/data/HelloWorld.bf someInput
node tmp/translated.js


go test -v ./...