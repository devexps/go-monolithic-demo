
# initialize develop environment
init:
	@which micro &> /dev/null || go install github.com/devexps/go-micro/cmd/micro/v2@latest
	@which protoc-gen-go-http &> /dev/null || go install github.com/devexps/go-micro/cmd/protoc-gen-go-http/v2@latest

	@which protoc-gen-go &> /dev/null || go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0
	@which protoc-gen-validate &> /dev/null || go install github.com/envoyproxy/protoc-gen-validate@v1.0.0
	@which protoc-gen-go-grpc &> /dev/null || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@which protoc-gen-openapi &> /dev/null || go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	@which wire &> /dev/null || go install github.com/google/wire/cmd/wire@latest

# generate protobuf api go code
api:
	@cd ../../../ && \
	buf generate

# generate OpenAPI v3 doc
openapi:
	@cd ../../../ && \
	buf generate --path api/admin/service/v1 --template api/admin/service/v1/buf.openapi.gen.yaml

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