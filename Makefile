SHELL := /bin/bash

.DEFAULT_GOAL := help
.DELETE_ON_ERROR:

ROOT := $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))
SOURCE_DIR := $(ROOT)/source
INSIGHTS_DIR := $(ROOT)/../insights
CORE_DIR := $(ROOT)/../core/source
TEST_DIR := $(ROOT)/e2e-tests
BUILD_DIR ?= $(ROOT)/build/linux
TEST_BUILD_DIR ?= $(ROOT)/build/tests

GO ?= go
GOFMT ?= gofmt
CLANG ?= clang
LLC ?= llc
PODMAN ?= podman
SUDO ?= $(shell if [ "$$EUID" -eq 0 ]; then echo ""; else echo sudo; fi)

GO_TAGS ?= scout
TARGET_VARIANTS ?= scout
TARGET_ARCHES ?= amd64 arm64
TEST_DISTRIBUTIONS ?= debian alpine archlinux fedora opensuse almalinux rockylinux cblmariner rpm
TEST_ARCHES ?= amd64
TEST_TIMEOUT ?= 10m
E2E_IMAGE ?= tholian-scout-e2e

.PHONY: help
help:
	@echo ""
	@echo "Usage: make <target>"
	@echo ""
	@echo "  help          Show this help"
	@echo "  build         Build all Linux CLI binaries"
	@echo "  check         Run formatting, vet, and unit tests"
	@echo "  build-tests   Build all end-to-end test binaries"
	@echo "  build-images  Build all end-to-end test containers"
	@echo "  test          Run end-to-end tests inside test containers"
	@echo "  clean         Remove generated binaries and test logs"

.PHONY: check-go-tools
check-go-tools:
	@command -v $(GO) >/dev/null 2>&1 || { echo "missing go: $(GO)"; exit 1; }
	@command -v $(GOFMT) >/dev/null 2>&1 || { echo "missing gofmt: $(GOFMT)"; exit 1; }

.PHONY: check-podman
check-podman:
	@command -v $(PODMAN) >/dev/null 2>&1 || { echo "missing podman: $(PODMAN)"; exit 1; }
	@$(PODMAN) info >/dev/null 2>&1 || { echo "podman is not usable: $(PODMAN)"; exit 1; }

.PHONY: check-insights
check-insights:
	@test -d "$(INSIGHTS_DIR)" && test -f "$(INSIGHTS_DIR)/go.mod" || { \
		echo "Missing Tholian Insights module:"; \
		echo ""; \
		echo "    $(INSIGHTS_DIR)/go.mod"; \
		echo ""; \
		echo "This directory is required for Tholian Scout builds and is not tracked by Git."; \
		echo "See README.md."; \
		echo ""; \
		exit 1; \
	}

.PHONY: check-core
check-core:
	@test -d "$(CORE_DIR)" && test -f "$(CORE_DIR)/go.mod" || { \
		echo "Missing core module:"; \
		echo ""; \
		echo "    $(CORE_DIR)/go.mod"; \
		echo ""; \
		echo "This directory is required for guard builds and is not tracked by Git."; \
		echo "See README.md."; \
		echo ""; \
		exit 1; \
	}

.PHONY: build
build: | check-go-tools check-insights check-core
	@mkdir -p "$(BUILD_DIR)"
	@for variant in $(TARGET_VARIANTS); do \
		for arch in $(TARGET_ARCHES); do \
			echo "Build $$variant / $$arch"; \
			CGO_ENABLED=0 GOOS=linux GOARCH=$$arch $(GO) -C "$(SOURCE_DIR)" build -tags "$$variant" -o "$(BUILD_DIR)/tholian-$$variant-$$arch" ./cmds/tholian-scout/main.go || exit 1; \
		done; \
	done

.PHONY: build-tests
build-tests: | check-go-tools check-insights check-core
	@mkdir -p "$(TEST_BUILD_DIR)"
	@for distro in $(TEST_DISTRIBUTIONS); do \
		for arch in $(TEST_ARCHES); do \
			echo "Build tests $$distro / $$arch"; \
			CGO_ENABLED=0 GOOS=linux GOARCH=$$arch $(GO) -C "$(TEST_DIR)" test -tags "$(GO_TAGS)" -c -o "$(TEST_BUILD_DIR)/$$distro-$$arch" "./$$distro" || exit 1; \
		done; \
	done

.PHONY: build-images
build-images: | check-podman
	@for distro in $(TEST_DISTRIBUTIONS); do \
		for arch in $(TEST_ARCHES); do \
			echo "Build test image $$distro / $$arch"; \
			$(PODMAN) build --platform linux/$$arch -f "$(TEST_DIR)/$$distro/Containerfile" -t "$(E2E_IMAGE):$$distro-$$arch" "$(ROOT)" || exit 1; \
		done; \
	done

.PHONY: check
check: | check-go-tools check-insights check-core
	@failed=0; \
	for dir in "$(SOURCE_DIR)" "$(TEST_DIR)"; do \
		files="$$(cd "$$dir" && $(GOFMT) -l . 2>&1)"; \
		if [ -n "$$files" ]; then \
			printf '%s\n' "$$files"; \
			failed=1; \
		fi; \
	done; \
	if [ $$failed -ne 0 ]; then \
		echo "gofmt: unformatted files"; \
		exit 1; \
	fi; \
	echo "gofmt: clean"
	$(GO) -C "$(SOURCE_DIR)" vet -tags "$(GO_TAGS)" ./...
	$(GO) -C "$(TEST_DIR)" vet -tags "$(GO_TAGS)" ./...
	$(GO) -C "$(SOURCE_DIR)" test -tags "$(GO_TAGS)" -race -cover ./...

.PHONY: test
test: | check-podman build-tests build-images
	@for distro in $(TEST_DISTRIBUTIONS); do \
		for arch in $(TEST_ARCHES); do \
			echo "Test $$distro / $$arch"; \
			$(PODMAN) run --rm --platform linux/$$arch \
				-v "$(TEST_BUILD_DIR)/$$distro-$$arch:/tests/e2e:ro" \
				"$(E2E_IMAGE):$$distro-$$arch" /tests/e2e -test.v -test.timeout "$(TEST_TIMEOUT)" || exit 1; \
		done; \
	done

.PHONY: clean-images
clean-images: | check-podman
	@for distro in $(TEST_DISTRIBUTIONS); do \
		for arch in $(TEST_ARCHES); do \
			$(PODMAN) rmi -f "$(E2E_IMAGE):$$distro-$$arch" >/dev/null 2>&1 || true; \
		done; \
	done

.PHONY: clean
clean:
	rm -rf "$(BUILD_DIR)" "$(TEST_BUILD_DIR)"
