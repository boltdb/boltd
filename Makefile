default: build

assets: assets/bindata.go

# go get github.com/jteeuwen/go-bindata/...
assets/bindata.go: assets/*.css assets/*.js
	@cd assets && go-bindata -pkg=assets -ignore=bindata.go .

build: assets ego
	mkdir -p bin
	go build -o bin/boltd ./cmd/boltd

# go get github.com/benbjohnson/ego/...
ego:
	@ego .

.PHONY: assets
