BOLD = \033[1m
CLEAR = \033[0m
CYAN = \033[36m

help: ## Display this help
	@awk '\
		BEGIN {FS = ":.*##"; printf "Usage: make $(CYAN)<target>$(CLEAR)\n"} \
		/^[a-z0-9]+([\/]%)?([\/](%-)?[a-z\-0-9%]+)*:.*? ##/ { printf "  $(CYAN)%-15s$(CLEAR) %s\n", $$1, $$2 } \
		/^##@/ { printf "\n$(BOLD)%s$(CLEAR)\n", substr($$0, 5) }' \
		$(MAKEFILE_LIST)

##@: Development

dev/setup: ## Setup local dependencies
	@go mod tidy

dev/up: ## Run development server(s)
	@go run github.com/mattn/goreman start

dev/dbmigrate-%: ENV=$*
dev/dbmigrate-%: ## Run database migrations
	@go run github.com/gobuffalo/pop/v6/soda create --env $(ENV) 2>/dev/null || true
	@go run github.com/gobuffalo/pop/v6/soda migrate --env $(ENV)

dev/migration: ## Generate a database migration named $(NAME)
	@go run github.com/gobuffalo/pop/v6/soda generate fizz "$(NAME)"

##@: Test

test: ## Run unit tests
	@go test -cover ./pkg/...
