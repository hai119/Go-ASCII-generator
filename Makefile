.PHONY: build test clean

build:
	go build -o bin/ascii ./cmd/ascii

test:
	go test -v ./...

clean:
	rm -rf bin/

run-image2text:
	./bin/ascii --mode image2text --input examples/input.jpg --output examples/output.txt

run-image2image:
	./bin/ascii --mode image2image --input examples/input.jpg --output examples/output.jpg 