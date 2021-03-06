LOCAL_PREFIX := github.com/jokeyrhyme/go-my-github-repos
MAIN_PKG := $(LOCAL_PREFIX)/cmd/my-github-repos

build: fmt
	go build -o ./build/bin/my-github-repos $(MAIN_PKG)

fmt:
	goreturns -b -l -local $(LOCAL_PREFIX) -w ./cmd/**/*.go ./pkg/**/*.go

lint: 
	gometalinter ./cmd/... ./pkg/...

setup:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
	go get -u github.com/golang/dep/cmd/dep
	dep ensure
	go get -u golang.org/x/tools/cmd/goimports
	go get -u sourcegraph.com/sqs/goreturns

test: fmt lint test-cover test-race

test-cover: fmt lint
	go test -cover ./cmd/... ./pkg/...

test-race: fmt lint
	go test -race ./cmd/... ./pkg/...

.PHONY:
