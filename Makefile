build: ## build
	@go build -o bin/ecom cmd/main.go

test: ## test
	@go test -v ./...

run: build ## run
	@./bin/ecom

setting-mysql: ## setting infra
	@docker-compose -f infra/local/docker-compose.yml up -d

down-mysql: ## down mysql docker container
	@docker-compose -f infra/local/docker-compose.yml down

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

migration: ## migration ex) make migration add-user-table
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up: ## migrate up
	@go run cmd/migrate/main.go up

migrate-down: ## migrate down
	@go run cmd/migrate/main.go down