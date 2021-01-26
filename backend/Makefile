ENV_FILE := .env
ENV = $(shell cat $(ENV_FILE))

.PHONY:serve
serve:
	$(ENV) go run .

.PHONY:test
test:
	go test -v ./... -count=1

.PHONY:start-deps
start-deps:
	docker-compose -f docker-compose.deps.yaml up -d

.PHONY:stop-deps
stop-deps:
	docker-compose -f docker-compose.deps.yaml down