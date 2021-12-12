.PHONY: test
test:
	@go clean -testcache
	@go test -parallel 1 -v

