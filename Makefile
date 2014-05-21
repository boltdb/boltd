default: build

build: templates
	mkdir -p bin
	go build -o bin/boltd ./cmd/boltd

# go get github.com/benbjohnson/ego/...
templates:
	@ego .

.PHONY: templates
