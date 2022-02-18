# brainfuck_to_js_translator
Translator of Brainfuck to Javascript

# Limitations and plans
1 - Use Abstract syntax tree to reduce JS command count
2 - Make input argument optional

# How to run
    set GOPATH=.
    go run src/main/main.go test/data/ReverseInput.bf test/data/ReverseInput.in
    node tmp/translated.js

# How to test
    go test -v .\test\translator-service-test\...

# How to test with external repo
    git submodule add https://github.com/rdebath/Brainfuck .\test\external_data\