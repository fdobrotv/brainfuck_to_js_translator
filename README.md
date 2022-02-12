# brainfuck_to_js_translator
Translator of Brainfuck to Javascript

set GOPATH=.
go run cmd/translator-service/main.go test/data/HelloWorld.bf someInput
node tmp/translated.js