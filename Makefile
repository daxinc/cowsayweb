# If you update this file, please follow
# https://suva.sh/posts/well-documented-makefiles

.DEFAULT_GOAL:=help

MAIN_PACKAGE := creditoffers
VERSION := 0.0.1
BUILT_ON := $(shell date)
COMMIT_HASH := $(shell git log -n 1 --pretty=format:"%H")
LDFLAGS := '-X "main.builtOn=$(BUILT_ON)" -X main.commitHash=$(COMMIT_HASH) -X main.version=$(VERSION)'
GO_LINUX := GOOS=linux GOARCH=amd64
GO_OSX := GOOS=darwin GOARCH=amd64
GO_WINDOWS := GOOS=windows GOARCH=386
GO_DOCKER := CGO_ENABLED=0 GOOS=linux

osx: ## Build MacOS executable
	$(GO_OSX) go build -o dist/cowsayweb -ldflags $(LDFLAGS) .

linux: ## Build Linux executable
	$(GO_LINUX) go build -o dist/cowsayweb-linux -ldflags $(LDFLAGS) .

windows: ## Build Windows executable
	$(GO_WINDOWS) go build -o dist/cowsayweb.exe -ldflags $(LDFLAGS) .

build: ## Build docker image
	$(GO_DOCKER) go build -mod vendor -a -installsuffix cgo -o cowsayweb -ldflags $(LDFLAGS) .
	docker build -t cowsayweb:latest .

run: ## Run the docker image
	docker run -d -p 8080:8080 cowsayweb

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[%.a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)