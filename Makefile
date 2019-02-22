SOURCE = $(wildcard *.go)
TAG ?= $(shell git describe --tags)
GOBUILD = go build -ldflags '-w'
TEST?=./...

ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")

setup:
	mkdir -p ${GOPATH}/bin
	go get -u golang.org/x/lint/golint
	go get -u github.com/kardianos/govendor
	go get github.com/mattn/goveralls

ALL = \
	$(foreach arch,x64 x32,\
	$(foreach suffix,linux osx windows,\
		build/charlatan-$(suffix)-$(arch))) \
	$(foreach arch,arm arm64,\
		build/charlatan-linux-$(arch))


help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

all: install test build

tools: ## Install tools
	go get -u golang.org/x/tools/cmd/stringer
	go get -u github.com/davecgh/go-spew/spew
	go get -u github.com/client9/misspell/cmd/misspell

build: clean test $(ALL) ## Build all binaries

spelling:
	misspell .

test: fmt vet
	go list $(TEST) | xargs -t -n4 go test $(TESTARGS) -v -timeout=60s -parallel=4

clean: ## Clean Target
	rm -f $(ALL)

lint:
	golint $(ALL_PACKAGES)

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

test-cover-html:
	@echo "mode: count" > coverage-all.out
	$(foreach pkg, $(ALL_PACKAGES),\
	ENVIRONMENT=test go test -coverprofile=coverage.out -covermode=count $(pkg);\
	tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out -o coverage.html

vendor-status:
	@govendor status

# os is determined as thus: if variable of suffix exists, it's taken, if not, then
# suffix itself is taken
osx = darwin
build/charlatan-%-x64: $(SOURCE)
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=$(firstword $($*) $*) GOARCH=amd64 $(GOBUILD) -o $@

build/charlatan-%-x32: $(SOURCE)
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=$(firstword $($*) $*) GOARCH=386 $(GOBUILD) -o $@

build/charlatan-linux-arm: $(SOURCE)
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) -o $@

build/charlatan-linux-arm64: $(SOURCE)
	@mkdir -p $(@D)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -o $@
