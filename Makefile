#
# go-gatest Makefile
#
# Examples:
#   make init
#   make clean
#   make build
#   make release
#

VERSION := $(shell cat VERSION)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_DIRTY := $(shell test -n "$$(git status --porcelain)" && echo +dev)
APP_VERSION := $(VERSION)-$(GIT_COMMIT)$(GIT_DIRTY)

BUILD_DIR ?= ./bin
GO ?= go
COMMON_FLAGS := -Og -g
COMMON_GO_FLAGS := -gcflags="all=-N -l" -buildmode=pie -tags='dev libsqlite3' -trimpath
COMMON_GO_LDFLAGS := -X main.version=$(APP_VERSION)

ifeq ($(shell uname -s),Darwin)
	COMMON_FLAGS += -fstack-clash-protection -fstack-protector-strong
	COMMON_LDFLAGS := -Wl,-dead_strip
else
	COMMON_FLAGS += -fno-plt -fstack-clash-protection -fstack-protector-strong -D_FORTIFY_SOURCE=3
	COMMON_LDFLAGS := -Wl,-z,relro,-z,now,-z,noexecstack
endif

define build_debug
	CGO_ENABLED=1 \
	CGO_CFLAGS="$(COMMON_FLAGS)" \
	CGO_CXXFLAGS="$(COMMON_FLAGS)" \
	CGO_FFLAGS="$(COMMON_FLAGS)" \
	CGO_LDFLAGS="$(COMMON_LDFLAGS)" \
	$(GO) build \
		$(COMMON_GO_FLAGS) \
		-ldflags '$(COMMON_GO_LDFLAGS)' \
		-o $(BUILD_DIR)/$(1) \
		./cmd/$(1)
endef

.PHONY: build clean init release tools goreleaser-check

goreleaser-check:
	@command -v goreleaser >/dev/null 2>&1 || (echo "No goreleaser command found. Install with: make tools (Linux) or brew install goreleaser (macOS)."; exit 1)
build:
	$(call build_debug,go-gatest)

clean:
	go clean

release: goreleaser-check
	goreleaser release --clean --snapshot

tools:
	go install github.com/goreleaser/goreleaser/v2@latest

init: tools
	go mod download
