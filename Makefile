deps:
	go get -u github.com/maxbrunsfeld/counterfeiter
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/onsi/ginkgo/ginkgo
	go get -u github.com/onsi/gomega

format:
	go get golang.org/x/tools/cmd/goimports
	find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +
	find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w "{}" +

test:
	@go test -cover -race $(shell go list ./... | grep -v /vendor/)
