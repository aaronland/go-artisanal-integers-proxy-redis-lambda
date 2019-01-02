prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/aaronland/go-artisanal-integers-proxy-redis
	# cp *.go src/github.com/aaronland/go-artisanal-integers-proxy-redis/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

deps:
	@GOPATH=$(shell pwd) go get "github.com/aaronland/go-artisanal-integers-proxy"
	@GOPATH=$(shell pwd) go get "github.com/whosonfirst/go-whosonfirst-pool-redis"
	mv src/github.com/aaronland/go-artisanal-integers-proxy/vendor/github.com/aaronland/* src/github.com/aaronland/
	mv src/github.com/aaronland/go-artisanal-integers-proxy/vendor/github.com/whosonfirst/* src/github.com/whosonfirst/
	rm -rf src/github.com/whosonfirst/go-whosonfirst-pool-redis/vendor/github.com/whosonfirst/go-whosonfirst-pool

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go

bin:	self
	@GOPATH=$(shell pwd) go build -o bin/proxy-server cmd/proxy-server.go

