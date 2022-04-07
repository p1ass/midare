ENV_FILE := .env
ENV_EXAMPLE_FILE := .env.example
ENV = $(shell cat $(ENV_FILE))
ENV_EXAMPLE = $(shell cat $(ENV_EXAMPLE_FILE))

.PHONY:serve
serve:
	$(ENV) go run .

.PHONY:test
test:
	$(ENV_EXAMPLE) go test -v ./... -count=1

.PHONY:start-deps
start-deps:
	docker-compose -f docker-compose.deps.yaml up -d

.PHONY:stop-deps
stop-deps:
	docker-compose -f docker-compose.deps.yaml down