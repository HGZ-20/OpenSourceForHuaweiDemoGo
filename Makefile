# 变量定义
APP = user
MAIN = user.go
OUTPUT = bin/$(APP)
PROTO_DIR = apps/user/rpc
API_DIR = api
API_OUTPUT = apps/user/api
RPC_OUTPUT = apps/user/rpc
BIN_DIR = bin
API_CONFIG = apps/user/api/etc/user.yaml
RPC_CONFIG = apps/user/rpc/etc/user.yaml

# 默认目标
.PHONY: all
all: build

# 构建目标
.PHONY: build
build: build-api build-rpc
	@echo "Build complete!"

# 构建 API 服务
.PHONY: build-api
build-api:
	@echo "Building API service..."
	go build -o $(BIN_DIR)/api_service $(API_OUTPUT)/$(APP).go

# 构建 RPC 服务
.PHONY: build-rpc
build-rpc:
	@echo "Building RPC service..."
	go build -o $(BIN_DIR)/rpc_service $(RPC_OUTPUT)/$(APP).go

# 运行目标
.PHONY: run
run: run-api run-rpc

# 运行 API 服务
.PHONY: run-api
run-api:
	@echo "Running API service..."
	$(BIN_DIR)/api_service -f $(API_CONFIG)

# 运行 RPC 服务
.PHONY: run-rpc
run-rpc:
	@echo "Running RPC service..."
	$(BIN_DIR)/rpc_service -f $(RPC_CONFIG)

# 生成 proto 文件
.PHONY: proto
proto:
	@echo "Generating gRPC and Go Zero proto files..."
	cd $(PROTO_DIR) && goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero

# 生成 API 文件
.PHONY: api
api:
	@echo "Generating Go Zero API files..."
	cd apps && goctl api go -api user/api/user.api -dir user/api -style gozero

# 清理目标
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf $(OUTPUT)

# 更新依赖
.PHONY: deps
deps:
	@echo "Updating dependencies..."
	go mod tidy

# 格式化代码
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 运行测试
.PHONY: test
test:
	@echo "Running tests..."
	go test ./... -v

