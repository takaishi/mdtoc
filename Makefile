TEST ?= $(shell go list ./... | grep -v -e vendor -e keys -e tmp)
VERSION = $(shell cat version)
BUILD=tmp/bin

INFO_COLOR=\033[1;34m
RESET=\033[0m
BOLD=\033[1m

default: build

depsdev: ## Installing dependencies for development
	go get github.com/golang/lint/golint
	go get -u github.com/tcnksm/ghr
	go get -u github.com/Songmu/goxz/cmd/goxz


test: ## Run test
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Testing$(RESET)"
	go test -v $(TEST) -timeout=30s -parallel=4
	go test -race $(TEST)

lint: ## Exec golint
	@echo "$(INFO_COLOR)==> $(RESET)$(BOLD)Linting$(RESET)"
	golint -min_confidence 1.1 -set_exit_status $(TEST)

build:
	go build  -ldflags "-X github.com/takaishi/mdtoc/config.Version=$(VERSION)" -o $(BUILD)/mdtoc

dist: clean
	goxz -pv=$(VERSION) -os=darwin,linux -arch=amd64 -d=dist -build-ldflags "-X github.com/takaishi/mdtoc/config.Version=$(VERSION)" .

github_release: ## Create some distribution packages
	ghr -u takaishi -r mdtoc --replace v$(VERSION) dist/

clean:
	rm -rf tmp
	rm -rf dist