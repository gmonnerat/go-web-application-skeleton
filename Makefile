all: bootstrap build

GO_BIN = go

bootstrap:
	$(GO_BIN) get -u github.com/tools/godep
	$(GO_BIN) get -u golang.org/x/sys/unix
	$(GO_BIN) get -u github.com/jteeuwen/go-bindata/...
	$(GO_BIN) get -u github.com/yosssi/goat/...
	$(GO_BIN) get -u github.com/pilu/fresh

build:
	go-bindata static/... templates/...
