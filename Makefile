BINARY=kratos
GOARCH=amd64
VERSION=0.0.1

lint:
	golint ./...

vet:
	go vet ./...

clean:
	rm -rf bin

install:
	go install -race -v

binary-depends:
	mkdir -p bin

# Builds
darwin-build: vet lint clean binary-depends
	env GOOS=darwin GOARCH=$(GOARCH) go build -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).darwin

dragonfly-build: vet lint clean binary-depends
	env GOOS=dragonfly GOARCH=$(GOARCH) go build -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).dragonfly

freebsd-build: vet lint clean binary-depends
	env GOOS=freebsd GOARCH=$(GOARCH) go build -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).freebsd

linux-build: vet lint clean binary-depends
	env GOOS=linux GOARCH=$(GOARCH) go build -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).linux

netbsd-build: vet lint clean binary-depends
	env GOOS=netbsd GOARCH=$(GOARCH) go build -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).netbsd

openbsd-build: vet lint clean binary-depends
	env GOOS=openbsd GOARCH=$(GOARCH) go build -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).openbsd

solaris-build: vet lint clean binary-depends
	env GOOS=solaris GOARCH=$(GOARCH) go build -v -o bin/$(BINARY).release.$(VERSION).$(GOARCH).solaris


build: darwin-build dragonfly-build freebsd-build linux-build netbsd-build openbsd-build solaris-build

.PHONY: vet install binary-depends lint
