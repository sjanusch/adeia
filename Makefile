deps:
	@go get -u github.com/maxbrunsfeld/counterfeiter
	@go get -u golang.org/x/tools/cmd/goimports
	@go get -u github.com/onsi/ginkgo/ginkgo
	@go get -u github.com/onsi/gomega

format:
	@go get golang.org/x/tools/cmd/goimports
	@find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +
	@find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w "{}" +

test:
	@go test -cover -race $(shell go list ./... | grep -v /vendor/)

# go quality checks
check: format lint vet

# vet entire repo (excluding vendor)
vet:
	@go vet $(shell go list ./... | grep -v /vendor/)

# lint entire repo (excluding vendor)
lint:
	@golint -min_confidence 1 $(shell go list ./... | grep -v /vendor/)

# errcheck entire repo (excluding vendor)
errcheck:
	@errcheck -ignore '(Close|Write)' $(shell go list ./... | grep -v /vendor/)

cover:
	go get github.com/haya14busa/goverage
	go get github.com/schrej/godacov
	goverage -v -coverprofile=coverage.out $(shell go list ./... | grep -v /vendor/)
