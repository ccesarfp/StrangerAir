all: dev-up

prod-up:
	@echo "Starting..."
	@docker-compose -f docker/prod/compose.yaml up

prod-down:
	@echo "Stopping..."
	@docker-compose -f docker/prod/compose.yaml down

dev-up:
	@echo "Starting dev environment..."
	@docker-compose -f docker/dev/compose.yaml up

dev-down:
	@echo "Stopping dev environment..."
	@docker-compose -f docker/dev/compose.yaml down

dev-access:
	@docker-compose -f docker/dev/compose.yaml exec -ti app /bin/sh

run-api:
	@go run cmd/api/*.go

watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

test:
	@go test ./... -v

clean:
	@rm -f main

.PHONY: all prod-up prod-down dev-up dev-down dev-access run-api test clean watch
