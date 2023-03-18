## Help
help:
	@printf "Available targets:\n\n"
	@awk '/^[a-zA-Z\-\_0-9%:\\]+/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
		helpCommand = $$1; \
		helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
	gsub("\\\\", "", helpCommand); \
	gsub(":+$$", "", helpCommand); \
		printf "  \x1b[32;01m%-35s\x1b[0m %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST) | sort -u
	@printf "\n"

default: help

## Start container
start:
	docker-compose up -d --build

## Stop
stop:
	docker-compose down

## Run unit test
test/unit:
	docker-compose -f ci.yaml build
	docker-compose -f ci.yaml run --rm cardgameapici go test

## Run static check
check:
	docker-compose -f ci.yaml build
	docker-compose -f ci.yaml run --rm cardgameapici staticcheck
