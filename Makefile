PROTO_FILE= deploy/school.proto
API_FILE= deploy/school.api
RPC_DIR= rpc
API_DIR= api
CODE_TEMPLATE= deploy/template
CODE_STYLE= go_zero

GOARCH=amd64
LDFLAGS := -s -w
PROJECT_NAME= primary-school-system

DOCKER_URL=registry.cn-hangzhou.aliyuncs.com/xxcheng
DOCKER_USERNAME=developer@xxcheng.cn

VERSION=$(shell git describe --tags --always)

.PHONY: gen-rpc
gen-rpc:
	@echo "Generating RPC code..."
	goctl rpc protoc $(PROTO_FILE) --go_out=$(RPC_DIR) --go-grpc_out=$(RPC_DIR) --zrpc_out=$(RPC_DIR) --style $(CODE_STYLE) --multiple

.PHONY: gen-api
gen-api:
	@echo "Generating API code..."
	goctl api go --api $(API_FILE) --dir $(API_DIR) --style $(CODE_STYLE) --home $(CODE_TEMPLATE)

.PHONY: gen-ent
gen-ent:
	@echo "Generating ent code..."
	go generate ./ent

.PHONY: run-rpc 
run-rpc:
	@echo "Running RPC server..."
	go run $(RPC_DIR)/school.go -f $(RPC_DIR)/etc/school.yaml

.PHONY: run-api
run-api:
	@echo "Running RPC server..."
	go run $(API_DIR)/school.go -f $(API_DIR)/etc/school.yaml

.PHONY: format-api
format-api:
	@echo "Starting Format Api file..."
	goctl api format --dir deploy


.PHONY: build-linux-rpc
build-linux-rpc:
	@echo "Build project of rpc for Linux..."
	env CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o ./build/$(PROJECT_NAME)_rpc ./rpc/school.go

.PHONY: build-linux-api
build-linux-api:
	@echo "Build project of api for Linux..."
	env CGO_ENABLED=0 GOOS=linux GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -trimpath -o ./build/$(PROJECT_NAME)_api ./api/school.go

.PHONY: build-linux
build-linux: build-linux-rpc build-linux-api