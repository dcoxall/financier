default: test

deps:
	@go list -f '{{ join .Imports "\n"}}' | xargs -n1 go get -d

testdeps: deps
	@go list -f '{{ join .TestImports "\n"}}' | xargs -n1 go get -d
	@go test -i ./...

test: testdeps
	@go test ${TESTARGS} ./...

.PHONY : deps testdeps test
