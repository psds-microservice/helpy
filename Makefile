.PHONY: help init build run run-dev migrate migrate-create worker test test-api test-db \
        version clean proto proto-all proto-clean proto-help lint vet fmt docker-build \
        docker-run docker-compose-up docker-compose-down install-deps health-check \
        deps generate-docs bench load-test security-check dev

PROTOC_IMAGE = local/protoc-go:latest
PROTO_ROOT = pkg/proto
GEN_DIR = pkg/gen

# Главная цель по умолчанию
.DEFAULT_GOAL := help

## 📦 Proto файлы
proto: proto-build proto-generate

proto-build:
	@echo "📦 Building protoc-go image..."
	docker build -t $(PROTOC_IMAGE) -f protoc-go.Dockerfile .
	@echo "✅ Docker image built"

proto-generate:
	@echo "🔧 Generating Go code from shared proto files..."
	docker run --rm \
		-v "$(CURDIR):/workspace" \
		$(PROTOC_IMAGE)
	@echo "✅ Proto files generated"

proto-clean:
	@echo "🧹 Cleaning generated files..."
	@if exist "pkg\gen" rmdir /s /q "pkg\gen" 2>nul || rm -rf pkg/gen
	@echo "✅ Clean complete"
