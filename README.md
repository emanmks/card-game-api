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
make test/unit
```

Run static checking:

```bash
make check
```

Stop container:

```bash
make stop
```

## Todo list

- [x] Create skeleton
- [ ] Consumer can see a list of available cards and their code
- [ ] Consumer can create a new deck
- [ ] Consumer can open a deck
- [ ] Consumer can request a draw for a specific deck
