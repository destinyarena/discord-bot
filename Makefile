.PHONY: all clean build install uninstall reinstall docker docker-build docker-push

GORUN = go run
GOBUILD = go build

all: build

clean:
	rm -rf bin

build: clean
	$(GOBUILD) -o bin/d2arena cmd/bot/main.go

install:
	install -m0755 bin/d2arena /usr/bin/d2arena

uninstall:
	rm -rf /usr/bin/d2arena

reinstall: uninstall install


docker-build:
	test $(DOCKERREPO)
	docker build . -t $(DOCKERREPO)

docker-push:
	test $(DOCKERREPO)
	docker push $(DOCKERREPO)


docker: docker-build docker-push
