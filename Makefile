GOVERSION=$(shell go version)
THIS_GOOS=$(word 1,$(subst /, ,$(lastword $(GOVERSION))))
THIS_GOARCH=$(word 2,$(subst /, ,$(lastword $(GOVERSION))))
GOOS?=$(THIS_GOOS)
GOARCH?=$(THIS_GOARCH)
DIR_BUILD=build
DIR_SCRIPT=./misc/ops
DIR_CACHE=$(DIR_BUILD)/.cache
APP=$(DIR_BUILD)/$(GOOS)_$(GOARCH)/sample

default: sample
all: sample

.PHONY: \
	all \
	sample \
	clean \
	default \
	test \

$(DIR_CACHE)/mod:
	mkdir -p $(DIR_CACHE)
	ln -sfn $(GOPATH)/pkg/mod $@

test:
	go test -v -cover -coverprofile=cover.out ./...
	go tool cover -func=cover.out

runner-test:
	go test ./...

sample: $(DIR_CACHE)/mod
	go build -v -o $(APP) cmd/server/*.go

clean:
	-rm -rf $(DIR_BUILD)