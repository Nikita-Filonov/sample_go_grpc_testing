# include .env
GREEN=\033[0;32m
NC=\033[0m
DIM=\033[2m

UNAME_OS := $(shell uname -s)
UNAME_ARCH := $(shell uname -m)

TMP_BASE := .tmp
TMP := $(TMP_BASE)/$(UNAME_OS)/$(UNAME_ARCH)
TMP_BIN = $(TMP)/bin

export GO111MODULE := on
export GOBIN := $(abspath $(TMP_BIN))
export PATH := $(GOBIN):$(PATH)

testing:
	@printf "  >  $(GREEN)Testing...$(NC)\n"

success:
	@printf "  >  $(GREEN)Success!$(NC)\n"

## test: runs all test
test:
	$(MAKE) testing
	@printf "$(DIM)"
	@GOBIN=$(GOBIN) ALLURE_RESULTS_PATH=. ENV=local go test ./tests -v -timeout 2m
	@printf "$(NC)"
	$(MAKE) success
