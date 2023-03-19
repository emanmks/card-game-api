# Card Game API

REST api for card games.

## Prerequisites

- Linux or Mac Machine
- Docker Enginer
- Docker Compose

## Code Structure

- `src/thegame`: **Business logic** where entities, interfaces, and the logic located.
- `src/repository`: **Data layer** where developer can implement a different type of database driver.
- `src/api`: **API** the api handlers to exposed endpoints to public and consume the business logic to deliver resources.

## Development how-to

Start the api container:

```bash
make start
```

Run unit test:

```bash
make ci/unit-test
```

Run static checking:

```bash
make ci/static-check
```

Stop container:

```bash
make stop
```

## Rules

- `main` is the main branch
- Do not force push to main branch, instead create a PR and make sure all checks are passed

## Live

The new version will be auto-deployed whenever merge request is merged to main branch.

Live version can be found here: [https://dry-bush-2591.fly.dev/health](https://dry-bush-2591.fly.dev/health)

## Todo list

- [x] Create skeleton
- [x] Consumer can see a list of available cards and their code
- [x] Consumer can create a new deck
- [x] Consumer can open a deck
- [x] Consumer can request a draw for a specific deck
