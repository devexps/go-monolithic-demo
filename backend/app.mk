GOPATH ?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif

APP_VERSION=$(shell git describe --tags --always)
APP_RELATIVE_PATH=$(shell a=`basename $$PWD` && cd .. && b=`basename $$PWD` && echo $$b/$$a)
APP_NAME=$(shell echo $(APP_RELATIVE_PATH) | sed -En "s/\//-/p")
APP_DOCKER_IMAGE=$(shell echo "$(APP_NAME):$(APP_VERSION)" | awk -F '@' '{print "go-monolithic-demo/" $$0}')

# generate protobuf api go code
api:
	@cd ../../../ && \
	buf generate

# generate OpenAPI v3 doc
openapi:
	@cd ../../../ && \
	buf generate --path api/admin/service/v1 --template api/admin/service/v1/buf.openapi.gen.yaml

# generate wire code
wire:
	@go run -mod=mod github.com/google/wire/cmd/wire ./cmd/server

# generate ent code
ent:
ifneq ("$(wildcard ./internal/data/ent)","")
	@go run -mod=mod entgo.io/ent/cmd/ent generate \
				--feature privacy \
				--feature sql/modifier \
				--feature entql \
				--feature sql/upsert \
				./internal/data/ent/schema
endif

# generate code
gen: ent wire api openapi

# run application
run: api openapi
	@go run ./cmd/server -conf ./configs

# run tests
test:
	@go test ./...

# run coverage tests
cover:
	@go test -v ./... -coverprofile=coverage.out

# build service app
app: api wire ent build

# build golang application
build:
ifeq ("$(wildcard ./bin/)","")
	mkdir bin
endif
	@go build -ldflags "-X main.Version=$(APP_VERSION)" -o ./bin/ ./...

# build docker image
docker:
	@docker build \
		-t $(APP_DOCKER_IMAGE) \
    	-f ../../../.docker/Dockerfile \
    	--build-arg APP_RELATIVE_PATH=$(APP_RELATIVE_PATH) \
    	--build-arg APP_VERSION=$(APP_VERSION) \
    	--build-arg GRPC_PORT=9090 \
    	--build-arg HTTP_PORT=8080 \
    	../../../

# clean build files
clean:
	@go clean
	@rm -f "coverage.out"

# run static analysis
vet:
	@go vet -json ./...

# run lint
lint:
	@golangci-lint run

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo " make [target]"
	@echo ""
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help