prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/aaronland/go-artisanal-integers-proxy-redis-lambda
	cp -r util src/github.com/aaronland/go-artisanal-integers-proxy-redis-lambda/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

deps:
	@GOPATH=$(shell pwd) go get "github.com/aaronland/go-artisanal-integers-proxy-redis/cmd"
	@GOPATH=$(shell pwd) go get "github.com/aaronland/go-artisanal-integers-lambda"
	@GOPATH=$(shell pwd) go get "github.com/aws/aws-lambda-go/lambda"
	rm -rf src/github.com/aaronland/go-artisanal-integers-lambda/vendor/github.com/aaronland
	rm -rf src/github.com/aaronland/go-artisanal-integers-lambda/vendor/github.com/aws
	mv src/github.com/aaronland/go-artisanal-integers-proxy-redis/vendor/github.com/aaronland/* src/github.com/aaronland/
	mv src/github.com/aaronland/go-artisanal-integers-proxy-redis/vendor/github.com/whosonfirst src/github.com/

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt util/*.go

bin:	self
	@GOPATH=$(shell pwd) go build -o bin/proxy-func cmd/proxy-func.go
	@GOPATH=$(shell pwd) go build -o bin/proxy-server cmd/proxy-server.go

