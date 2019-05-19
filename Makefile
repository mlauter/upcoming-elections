all: build

deps:
	$(prefix)/bin/go get -u golang.org/x/lint/golint
	$(prefix)/bin/go get honnef.co/go/tools/cmd/megacheck
	$(prefix)/bin/go get github.com/gorilla/mux
	$(prefix)/bin/go get github.com/gorilla/handlers

lint: deps
	$(prefix)/bin/go fmt
	$(prefix)/bin/go vet
	$(GOPATH)/bin/golint -set_exit_status
	$(GOPATH)/bin/megacheck -unused.exit-non-zero -staticcheck.exit-non-zero

test: lint
	$(prefix)/bin/go test

build: clean test
	$(prefix)/bin/go build -o elections

clean:
	@rm -f elections

.PHONY: all deps lint test build
