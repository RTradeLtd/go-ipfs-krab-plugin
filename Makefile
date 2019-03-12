GO=env GO111MODULE=on go

.PHONY: build-plugin
build-plugin:
	mkdir build
	go build -o build/go-ipfs-krab-plugin.so --buildmode=plugin

.PHONY: clean
clean:
	rm -rf build

.PHONY: vet
vet:
	go vet ./...

.PHONY: deps
deps:
	$(GO) mod vendor
	$(GO) mod tidy