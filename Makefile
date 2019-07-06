ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

BIN_DIR=bin
TARGET=main.go
BINARY=jwtbin
VERSION=0.1.0
BUILD=`git rev-parse HEAD`
PLATFORMS=darwin linux windows
ARCHITECTURES=386 amd64

# Setup linker flags option for build that interoperate with variable names in src code
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

default: build

all: clean deps build

build:
	go build ${LDFLAGS} -o ${BIN_DIR}/${BINARY} ${TARGET}

build-all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES),\
	$(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build -o $(BIN_DIR)/$(BINARY)-$(GOOS)-$(GOARCH) $(TARGET))))
	@echo "Build Complete"

install:
	go install ${LDFLAGS}

deps:
	go get github.com/dgrijalva/jwt-go

# Remove only what we've created
clean:
	rm -rf "${ROOT_DIR}/${BIN_DIR}/${BINARY}*"

.PHONY: check clean install build-all all
