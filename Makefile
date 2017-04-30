ifndef GOPATH
$(warning You need to set up a GOPATH. Run "go help gopath".)
endif

# Just in case this is on an old version of go
export GO15VENDOREXPERIMENT := 1

GO:=$(strip $(shell which go 2> /dev/null))
PACKAGE  = github.com/wwsean08/go-action-sdk
DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
BIN      = $(GOPATH)/bin
BASE     = $(GOPATH)/src/$(PACKAGE)
packages =$(shell $(GO) list ./... | grep -v -e /vendor/ -e /sample)

.PHONY: clean test fmt vendor cover cover-ci

default: test

clean:
	rm -rf dist bin vendor

fmt:
	$(GO) fmt $(packages)

vendor:
	rm -rf vendor Godeps
	$(GO) get -insecure -v -t $(packages); godep save ./...; godep save $(packages)

test: fmt
	$(GO) vet $(packages)
	$(GO) test -race -cover $(packages)

cover:
	echo "mode: set" > coverage-all.out
	$(foreach pkg,$(packages),\
		go test -coverprofile=coverage.out $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -func=coverage-all.out

cover-ci: cover
	goveralls -coverprofile=coverage-all.out -service=travis-ci