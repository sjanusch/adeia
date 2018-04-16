# test entire repo
test:
	@go test -cover -race $(shell go list ./... | grep -v /vendor/)
