.PHONY: all clean build install uninstall reinstall

GORUN = go run
GOBUILD = go build

all: clean build

clean:
	rm -rf bin

build: clean
	$(GOBUILD) -o bin/d2arena cmd/bot/main.go

install:
	install -m0755 bin/d2arena /usr/bin/d2arena

uninstall:
	rm -rf /usr/bin/d2arena

reinstall: uninstall install
