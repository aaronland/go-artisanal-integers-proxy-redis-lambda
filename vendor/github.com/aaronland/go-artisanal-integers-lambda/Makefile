prep:
	if test -d pkg; then rm -rf pkg; fi
 
self:   prep
	if test -d src/github.com/aaronland/go-artisanal-integers-lambda; then rm -rf src/github.com/aaronland/go-artisanal-integers-lambda; fi
	mkdir -p src/github.com/aaronland/go-artisanal-integers-lambda/
	cp *.go src/github.com/aaronland/go-artisanal-integers-lambda/
	cp -r client src/github.com/aaronland/go-artisanal-integers-lambda/
	cp -r server src/github.com/aaronland/go-artisanal-integers-lambda/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

deps:
	@GOPATH=$(shell pwd) go get "github.com/aaronland/go-artisanal-integers"
	@GOPATH=$(shell pwd) go get "github.com/whosonfirst/algnhsa"
	@GOPATH=$(shell pwd) go get "github.com/aws/aws-sdk-go/aws"

vendor-deps: rmdeps deps
	if test ! -d src; then mkdir src; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt client/*.go
	go fmt cmd/*.go
	go fmt server/*.go

bin:    self
	if test ! -d bin; then mkdir bin; fi
	@GOPATH=$(shell pwd) go build -o bin/int cmd/int.go
