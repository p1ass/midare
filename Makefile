ENV_FILE := .env
ENV = $(shell cat $(ENV_FILE))

.PHONY:serve
serve:
	$(ENV) go run .

.PHONY:test
test:
	$(ENV) go test -v ./... -count=1

