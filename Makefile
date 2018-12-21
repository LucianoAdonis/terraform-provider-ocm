PROVIDER_NAME = ocm
VERSION = 0.0.3
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
TEST?=./...

BIN_DIRECTORY = _bin
EXECUTABLE_NAME = terraform-provider-$(PROVIDER_NAME)
DIST_ZIP_PREFIX = $(EXECUTABLE_NAME).v$(VERSION)




.PHONY: all fmt meh apply build restart debug destroy restart package help

fmt:
	gofmt -w $(GOFMT_FILES)

all: help

first-run: package build ## Downloads all packages and compiles the provider

meh: restart build ## Just Meh

apply: ## init & apply
	@terraform init
	@terraform apply

build: ## Make a binary for Terraform
	@go build -o $(EXECUTABLE_NAME)

debug:
	export TF_LOG=DEBUG
	export TF_LOG_PATH=./terraform.log

destroy: ## Destroy
	@terraform destroy

restart: ## Clean
	@rm terraform.tfs*
	@rm -Rf .terraform
	@rm terraform-provider-*

package: ## Get packages
	@dep init

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

