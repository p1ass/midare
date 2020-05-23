ENV_FILE := .env
ENV = $(shell cat $(ENV_FILE))

.PHONY:serve
serve:
	$(ENV) go run .

.PHONY:test
test:
	$(ENV) go test -v ./... -count=1

.PHONY: integration-test
integration-test:
	$(ENV) go test -v ./... -tags integration -count=1

.PHONY: deploy-dev
deploy-dev:
	$(ENV) docker-compose -f docker/docker-compose.deps.base.yml -f docker/docker-compose.deps.dev.yml -p dev up -d
	$(ENV) ENV_FILE=$(ENV_FILE) docker-compose -f docker/docker-compose.base.yml -f docker/docker-compose.dev.yml -p dev stop
	$(ENV) ENV_FILE=$(ENV_FILE) docker-compose -f docker/docker-compose.base.yml -f docker/docker-compose.dev.yml -p dev up -d
