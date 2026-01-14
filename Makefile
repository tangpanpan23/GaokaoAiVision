.PHONY: all build clean test docker prod prod_api prod_spider prod_score prod_ai prod_user prod_model prod_command

# 版本信息
VERSION := $(shell git describe --tags --always --dirty)
BUILD_TIME := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
GIT_COMMIT := $(shell git rev-parse --short HEAD)

# Go 参数
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
CGO_ENABLED := 0

# 构建参数
LDFLAGS := -X main.version=$(VERSION) -X main.buildTime=$(BUILD_TIME) -X main.gitCommit=$(GIT_COMMIT) -w -s

# 构建所有服务
prod: prod_api prod_spider prod_score prod_ai prod_user prod_model prod_command

# API服务
prod_api:
	@echo "Building API service..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o bin/api ./app/api
	@echo "API service built successfully"

# 爬虫服务
prod_spider:
	@echo "Building Spider RPC service..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o bin/spider-rpc ./app/spider/rpc
	@echo "Spider RPC service built successfully"

# 数据查询服务
prod_score:
	@echo "Building Score RPC service..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o bin/score-rpc ./app/score/rpc
	@echo "Score RPC service built successfully"

# AI分析服务
prod_ai:
	@echo "Building AI RPC service..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o bin/ai-rpc ./app/ai/rpc
	@echo "AI RPC service built successfully"

# 用户服务
prod_user:
	@echo "Building User RPC service..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o bin/user-rpc ./app/user/rpc
	@echo "User RPC service built successfully"

# 模型服务
prod_model:
	@echo "Building Model RPC service..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o bin/model-rpc ./app/model
	@echo "Model RPC service built successfully"

# 命令行工具
prod_command:
	@echo "Building Command tool..."
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "$(LDFLAGS)" -o bin/command ./app/command
	@echo "Command tool built successfully"

# 开发环境构建
dev:
	@echo "Building all services for development..."
	go build -o bin/api ./app/api
	go build -o bin/spider-rpc ./app/spider/rpc
	go build -o bin/score-rpc ./app/score/rpc
	go build -o bin/ai-rpc ./app/ai/rpc
	go build -o bin/user-rpc ./app/user/rpc
	go build -o bin/model-rpc ./app/model
	go build -o bin/command ./app/command
	@echo "All services built successfully"

# 代码生成
gen:
	@echo "Generating API code..."
	goctl api go -api ./app/api/desc/api.api -dir ./app/api/
	@echo "Generating Spider RPC code..."
	goctl rpc protoc app/spider/rpc/spider.proto --go-grpc_out=./pb --go_out=./pb --zrpc_out=. -style goZero --client=true
	@echo "Generating Score RPC code..."
	goctl rpc protoc app/score/rpc/score.proto --go-grpc_out=./pb --go_out=./pb --zrpc_out=. -style goZero --client=true
	@echo "Generating AI RPC code..."
	goctl rpc protoc app/ai/rpc/ai.proto --go-grpc_out=./pb --go_out=./pb --zrpc_out=. -style goZero --client=true
	@echo "Generating User RPC code..."
	goctl rpc protoc app/user/rpc/user.proto --go-grpc_out=./pb --go_out=./pb --zrpc_out=. -style goZero --client=true
	@echo "Code generation completed"

# 数据库迁移
migrate-up:
	@echo "Running database migrations..."
	go run app/command migrate up

migrate-down:
	@echo "Rolling back database migrations..."
	go run app/command migrate down

migrate-create:
	@echo "Creating new migration..."
	@read -p "Enter migration name: " name; \
	go run app/command migrate create $$name

# Docker 构建
docker-build:
	@echo "Building Docker images..."
	docker build -t lighthouse-volunteer-api:$(VERSION) -f Dockerfile.api .
	docker build -t lighthouse-volunteer-spider:$(VERSION) -f Dockerfile.spider .
	docker build -t lighthouse-volunteer-score:$(VERSION) -f Dockerfile.score .
	docker build -t lighthouse-volunteer-ai:$(VERSION) -f Dockerfile.ai .
	docker build -t lighthouse-volunteer-user:$(VERSION) -f Dockerfile.user .
	docker build -t lighthouse-volunteer-model:$(VERSION) -f Dockerfile.model .
	docker build -t lighthouse-volunteer-command:$(VERSION) -f Dockerfile.command .

# 清理
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -rf pb/
	@echo "Clean completed"

# 测试
test:
	@echo "Running tests..."
	go test ./...

# 依赖管理
tidy:
	@echo "Tidying go modules..."
	go mod tidy

# 格式化
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 代码检查
vet:
	@echo "Running go vet..."
	go vet ./...

# 预提交检查
pre-commit: fmt vet test

# 帮助
help:
	@echo "Available commands:"
	@echo "  prod               - Build all services for production"
	@echo "  prod_api           - Build API service"
	@echo "  prod_spider        - Build Spider RPC service"
	@echo "  prod_score         - Build Score RPC service"
	@echo "  prod_ai            - Build AI RPC service"
	@echo "  prod_user          - Build User RPC service"
	@echo "  prod_model         - Build Model RPC service"
	@echo "  prod_command       - Build Command tool"
	@echo "  dev                - Build all services for development"
	@echo "  gen                - Generate code from proto and api files"
	@echo "  migrate-up         - Run database migrations"
	@echo "  migrate-down       - Rollback database migrations"
	@echo "  migrate-create     - Create new migration"
	@echo "  docker-build       - Build Docker images"
	@echo "  clean              - Clean build artifacts"
	@echo "  test               - Run tests"
	@echo "  tidy               - Tidy go modules"
	@echo "  fmt                - Format code"
	@echo "  vet                - Run go vet"
	@echo "  pre-commit         - Run pre-commit checks"
	@echo "  help               - Show this help"
