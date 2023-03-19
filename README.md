# Card Game API

REST api for card games.

## Prerequisites

- Linux or Mac Machine
- Docker Enginer
- Docker Compose

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

## Live

Live version can be found here: [https://dry-bush-2591.fly.dev/health](https://dry-bush-2591.fly.dev/health)

## Todo list

- [x] Create skeleton
- [x] Consumer can see a list of available cards and their code
- [x] Consumer can create a new deck
- [x] Consumer can open a deck
- [ ] Consumer can request a draw for a specific deck
