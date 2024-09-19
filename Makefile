devFile = "compose.dev.yaml"

all: dev-up

prod-up:
	@echo "Starting..."
	@docker-compose up

prod-down:
	@echo "Stopping..."
	@docker-compose down

dev-up:
	@echo "Starting dev environment..."
	@docker-compose -f ${devFile} up

dev-down:
	@echo "Stopping dev environment..."
	@docker-compose -f ${devFile} down

dev-access:
	@docker-compose -f ${devFile} exec -ti app /bin/sh

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
