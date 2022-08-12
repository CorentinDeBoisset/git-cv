# If version is undefined, we use "NoVersion as a default"
ifeq ($(VERSION),)
	VERSION = NoVersion
endif

PACKAGE_NAME := $(shell go list -m -f '{{.Path}}')
BIN_DIR := bin
ENTRYPOINT := cmd/git-cv
GO_FILES := $(shell find . -type f -name '*.go')

BUILD_FLAGS := -ldflags "-extldflags -static -X main.ExecutableName=${EXECUTABLE} -X main.Version=${VERSION}"

EXECUTABLE := git-cv

.PHONY: all
all: ${BIN_DIR}/${EXECUTABLE}

.PHONY: clean
clean:
	rm -f ${BIN_DIR}/${EXECUTABLE}

# Build the executable
${BIN_DIR}/${EXECUTABLE}: ${GO_FILES}
	CGO_ENABLED=0 go build ${BUILD_FLAGS} -o $@ ./${ENTRYPOINT}

.PHONY: install
install:
	CGO_ENABLED=0 go install ${BUILD_FLAGS} ./${ENTRYPOINT}
