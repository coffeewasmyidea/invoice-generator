NAME = invoice-generator
REMOTE = github.com/coffeewasmyidea

DIRS := \
	. \

uniq = $(if $1,$(firstword $1) $(call uniq,$(filter-out $(firstword $1),$1)))
gofiles = $(foreach d,$(1),$(wildcard $(d)/*.go))
fmt = $(addprefix fmt-,$(1))
outsuffix = bin/$(NAME)

all: bin-linux-amd64 bin-linux-arm bin-windows-amd64 bin-windows-arm bin-darwin-amd64 bin-darwin-arm64

sha = $(shell git rev-parse --short HEAD || cat SHA | tr -d ' \n')

ifeq ($(VERSION),)
VERSION = $(shell git describe --tags --match 'v*' | tr -d 'v \n')
realv = $(shell printf $(VERSION) | cut -d'-' -f1)
ifneq ($(VERSION),$(realv))
commits = $(shell printf $(VERSION) | cut -d'-' -f2)
VERSION := $(realv).$(commits).$(sha)
endif
endif

dirty = $(shell git diff --shortstat 2> /dev/null | tail -n1 | tr -d ' \n')
ifneq ($(dirty),)
VERSION := $(VERSION).dev
endif

id = $(shell head -c20 /dev/urandom|od -An -tx1|tr -d ' \n')

bin-linux-amd64: $(outsuffix).linux-amd64
bin-linux-arm: $(outsuffix).linux-arm
bin-windows-amd64: $(outsuffix).windows-amd64.exe
bin-windows-arm: $(outsuffix).windows-arm.exe
bin-darwin-amd64: $(outsuffix).darwin-amd64
bin-darwin-arm64: $(outsuffix).darwin-arm64

ldflags = -ldflags '-s -w -extldflags "-static" -X "main.ver=$(VERSION)" -X "main.sha=$(sha)" -B 0x$(id)'

install: $(call gofiles,$(DIRS))
	go install -ldflags '-s -w -X "main.ver=$(VERSION)" -X "main.sha=$(sha)" -B 0x$(id)' .

$(outsuffix).linux-amd64: $(call gofiles,$(DIRS))
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -x -v $(ldflags) -o $@ .

$(outsuffix).linux-arm: $(call gofiles,$(DIRS))
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build $(ldflags) -o $@ .

$(outsuffix).windows-amd64.exe: $(call gofiles,$(DIRS))
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(ldflags) -o $@ .

$(outsuffix).windows-arm.exe: $(call gofiles,$(DIRS))
	CGO_ENABLED=0 GOOS=windows GOARCH=arm GOARM=7 go build $(ldflags) -o $@ .

$(outsuffix).darwin-amd64: $(call gofiles,$(DIRS))
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(ldflags) -o $@ .

$(outsuffix).darwin-arm64: $(call gofiles,$(DIRS))
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build $(ldflags) -o $@ .

clean:
	rm -f bin/*

.PHONY: clean all \
	bin-linux-amd64 bin-linux-arm bin-windows-amd64 bin-windows-arm bin-darwin-amd64 bin-darwin-arm64

.PHONY: version
version:
	@echo $(VERSION)
