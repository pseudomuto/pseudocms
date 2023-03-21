BOLD = \033[1m
CLEAR = \033[0m
CYAN = \033[36m

OS = $(shell uname -s)
ARCH = $(shell uname -m)

help: ## Display this help
	@awk '\
		BEGIN {FS = ":.*##"; printf "Usage: make $(CYAN)<target>$(CLEAR)\n"} \
		/^[a-z0-9]+([\/]%)?([\/](%-)?[a-z\-0-9%]+)*:.*? ##/ { printf "  $(CYAN)%-15s$(CLEAR) %s\n", $$1, $$2 } \
		/^##@/ { printf "\n$(BOLD)%s$(CLEAR)\n", substr($$0, 5) }' \
		$(MAKEFILE_LIST)

##@: Development

dev/setup: bin/buf ## Setup local dependencies
	@go mod tidy
	@bin/buf mod update

dev/up: ## Run development server(s)
	@go run github.com/mattn/goreman start

db/migrate-%: ENV=$*
db/migrate-%: ## Run database migrations
	@go run github.com/gobuffalo/pop/v6/soda create --env $(ENV) 2>/dev/null || true
	@go run github.com/gobuffalo/pop/v6/soda migrate --env $(ENV)

db/migration: ## Generate a database migration named $(NAME)
	@go run github.com/gobuffalo/pop/v6/soda generate fizz "$(NAME)"

##@: Code Gen and Formatting

generate: bin/buf bin/protoc-gen-go bin/protoc-gen-go-grpc ## Generate protos files and such
	@PATH=./bin buf generate

lint: bin/buf ## Lints the code base
	@bin/buf lint

format: bin/buf ## Format proto files
	@bin/buf format -w

##@: Test

test: ## Run unit tests
	@go test -cover ./pkg/...

##@: Binaries

bin/buf: ## Install buf to ./bin
	@mkdir -p ./bin
	@curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/v1.15.1/buf-$(OS)-$(ARCH)" \
		-o ./bin/buf && chmod +x ./bin/buf

bin/ctl: ## Build ctl to bin/pseudoctl
	@go build -o bin/pseudoctl ./cmd/ctl

bin/protoc-gen-go: ## Install protoc-gen-go
	@GOBIN=$(abspath ./bin) go install google.golang.org/protobuf/cmd/protoc-gen-go

bin/protoc-gen-go-grpc: ## Install protoc-gen-go-grpc
	@GOBIN=$(abspath ./bin) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
