ifneq (,$(wildcard ./.env))
	include .env
	export
endif

doctor:
	@echo "⏳ Checking dependencies..."
	@make check_binary BIN_NAME=go
	@make check_binary BIN_NAME=migrate
	@make check_binary BIN_NAME=docker
	@make check_binary BIN_NAME=mockery
	@make check_command COMMAND="docker ps"
	@echo "✅ All good, you are ready to go!"

check_binary:
	@which $(BIN_NAME) >/dev/null 2>&1 \
	&& echo "✅ \`$(BIN_NAME)\` installed" \
	|| { echo >&2 "❌ \`$(BIN_NAME)\` not installed"; exit 1; }

check_command:
	@$(COMMAND) >/dev/null 2>&1 \
	&& echo "✅ \`$(COMMAND)\` command execution successful" \
	|| { echo >&2 "❌ \`$(COMMAND)\` command execution unsuccessful"; $(COMMAND) >&2; exit 1; }

migrate:
	@migrate -source file://migrations \
	-database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" \
	up

backfill:
	@go run cmd/cli/backfill/backfill.go

gen:
	@go run cmd/cli/generate/generate.go

compose:
	@docker compose up

env:
	@cp .env.base .env
	@echo "✅ \`.env\` file generated"

build:
	@echo "⏳ Building..."
	@BUILD_START=$$(date +%s) \
	;CGO_ENABLED=0 go build -o ./bin/app ./cmd/app/main.go \
	&& echo "✅ Build took $$(($$(date +%s)-BUILD_START)) seconds"

run: build
	@./bin/app

test:
	@go test -v ./...

swag:
	@`go env GOPATH`/bin/swag init -g ../../cmd/app/main.go -d ./internal/api --parseInternal --parseDependency

mock:
	@mockery --recursive --dir=internal/domain/repository --name="^." --with-expecter --inpackage
