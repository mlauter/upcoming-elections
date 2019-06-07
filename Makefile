GOCMD=go

all: build

deps:
	$(GOCMD) get -u golang.org/x/lint/golint
	$(GOCMD) get honnef.co/go/tools/cmd/megacheck
	$(GOCMD) get github.com/gorilla/mux
	$(GOCMD) get github.com/gorilla/handlers
	$(GOCMD) get github.com/gorilla/schema
	$(GOCMD) get github.com/gorilla/csrf
	$(GOCMD) get google.golang.org/api/civicinfo/v2
	$(GOCMD) get google.golang.org/api/option

lint: deps
	$(GOCMD) fmt
	$(GOCMD) vet
	$(GOPATH)/bin/golint -set_exit_status
	$(GOPATH)/bin/megacheck -unused.exit-non-zero -staticcheck.exit-non-zero

test: lint
	$(GOCMD) test

build: clean test
	$(GOCMD) build -o elections

clean:
	@rm -f elections

.PHONY: all deps lint test build
