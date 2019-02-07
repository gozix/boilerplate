# Package dir
GO_DIR ?= $(shell pwd)

# Package import path
GO_PKG ?= $(shell go list -e -f "{{ .ImportPath }}")

# Package version
GO_VER ?= $(shell date -u +%Y-%m-%d.%H:%M:%S)

# Package binary prefix
GO_BIN ?= $(shell basename $$(dirname $(GO_PKG)))

# Package binnary delimiter
GO_DEL ?= -

# Set the mode for code-coverage
GO_TEST_COVERAGE_MODE ?= count
GO_TEST_COVERAGE_FILE_NAME ?= coverage.out

# Set a default `min_confidence` value for `golint`
GO_LINT_MIN_CONFIDENCE ?= 0.2

all: install-deps build compress

.PHONY: build
build:
	@echo "Build binaries"
	@ERR=0; \
	for CMD in $$(find "./cmd" -maxdepth 1 -mindepth 1 -type d -print); do \
		BIN=$$(basename "$${CMD}"); \
		go build -gcflags="-trimpath=$(GO_DIR)" -asmflags="-trimpath=$(GO_DIR)" \
				 -ldflags "-X $(GO_PKG)/cli/command.appVersion=$(GO_VER)" \
				 -o "$(GO_DIR)/.bin/$${BIN}/$(GO_BIN)$(GO_DEL)$${BIN}" "$${CMD}" || { \
			ERR=$$?; \
			break; \
		}; \
	done; \
	if [ $$ERR != 0 ]; then \
		exit $$ERR; \
	fi

.PHONY: generate
generate:
	@echo "Run easyjson"
	@go list -f "{{ .Dir }}" ./... | xargs -I "{}" grep -lrw "{}" -e "easyjson:json" | sort | uniq | xargs -I "{}" easyjson "{}" || exit 1

.PHONY: compress
compress:
	@echo "Compress binaries"
	@ERR=0; \
    for CMD in $$(find "./cmd" -maxdepth 1 -mindepth 1 -type d -print); do \
    	BIN=$$(basename "$${CMD}"); \
    	upx --best -q "$(GO_DIR)/.bin/$${BIN}/$(GO_BIN)$(GO_DEL)$${BIN}" || { \
    		ERR=$$?; \
    		break; \
    	}; \
    done; \
    if [ $$ERR != 0 ]; then \
    	exit $$ERR; \
    fi;

.PHONY: install-deps
install-deps:
	@echo "Install dependencies"
	@dep ensure -v -vendor-only

.PHONY: install-deps-dev
install-deps-dev:
	@echo "Install Dep"
	@go get github.com/golang/dep/cmd/dep

.PHONY: update-deps
update-deps:
	@echo "Update dependencies"
	dep ensure -v -update

.PHONY: update-deps-dev
update-deps-dev:
	@echo "Update Dep"
	@go get -u github.com/golang/dep/cmd/dep

.PHONY: test
test:
	@echo "Run unit tests"
	@go test -v ./...

.PHONY: test-with-coverage
test-with-coverage:
	@echo "Run unit tests with coverage"
	@go test -cover ./...

.PHONY: test-with-coverage-profile
test-with-coverage-profile:
	@echo "Run unit tests with coverage profile"
	@echo "mode: ${GO_TEST_COVERAGE_MODE}" > "${GO_TEST_COVERAGE_FILE_NAME}"
	@go test -coverpkg=`go list ./... | grep -vE 'command|domain' | tr '\n' ','` -covermode ${GO_TEST_COVERAGE_MODE} -coverprofile=${GO_TEST_COVERAGE_FILE_NAME} ./...
	@echo "Generate coverage report";
	@go tool cover -func="${GO_TEST_COVERAGE_FILE_NAME}";
	@rm "${GO_TEST_COVERAGE_FILE_NAME}";

.PHONY: fix
fix: fix-format fix-import

.PHONY: fix-import
fix-import:
	@echo "Fix imports"
	@errors=$$(goimports -l -w -local $(GO_PKG) $$(go list -f "{{ .Dir }}" ./...)); if [ "$${errors}" != "" ]; then echo "$${errors}"; fi

.PHONY: fix-format
fix-format:
	@echo "Fix formatting"
	@gofmt -w ${GO_FMT_FLAGS} $$(go list -f "{{ .Dir }}" ./...); if [ "$${errors}" != "" ]; then echo "$${errors}"; fi

.PHONY: lint
lint: lint-format lint-import lint-style

.PHONY: lint-format
lint-format:
	@echo "Check formatting"
	@errors=$$(gofmt -l ${GO_FMT_FLAGS} $$(go list -f "{{ .Dir }}" ./...)); if [ "$${errors}" != "" ]; then echo "$${errors}"; exit 1; fi

.PHONY: lint-import
lint-import:
	@echo "Check imports"
	@errors=$$(goimports -l -local $(GO_PKG) $$(go list -f "{{ .Dir }}" ./...)); if [ "$${errors}" != "" ]; then echo "$${errors}"; exit 1; fi

.PHONY: lint-style
lint-style:
	@echo "Check code style"
	@errors=$$(golint -min_confidence=${GO_LINT_MIN_CONFIDENCE} $$(go list ./...)); if [ "$${errors}" != "" ]; then echo "$${errors}"; exit 1; fi

.PHONY: clean
clean:
	@echo "Cleanup"
	@find . -type f -name "*easyjson*" -delete
	@find . -type f -name "*coverage*.out" -delete
