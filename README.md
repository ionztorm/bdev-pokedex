# bdev-pokedex

A command-line Pokédex application written in Go. This project is part of the [Boot.dev](https://www.boot.dev) Go course and demonstrates key Go concepts such as modularity, testing, and API interaction.

## Features

- Explore Pokémon regions and locations
- Inspect Pokémon details
- Catch Pokémon and manage your Pokédex
- Caching for efficient API usage
- REPL (Read-Eval-Print Loop) interface for interactive commands
- Modular code structure with unit tests

## Project Structure

```
.
├── go.mod
├── main.go
├── README.md
└── internal/
    ├── api/
    │   ├── constants.go
    │   └── fetch.go
    ├── command/
    │   ├── catch.go
    │   ├── command-utils.go
    │   ├── exit.go
    │   ├── explore.go
    │   ├── help.go
    │   ├── inspect.go
    │   ├── map.go
    │   ├── mapb.go
    │   ├── pokedex.go
    │   └── structs.go
    ├── pokecache/
    │   ├── cache.go
    │   └── cache_test.go
    ├── repl/
    │   ├── repl.go
    │   └── repl_test.go
    └── utils/
       └── utils.go
```

## Getting Started

### Prerequisites

- Go 1.18 or later

### Installation

1. Clone the repository:
   ```sh
   git clone <repo-url>
   cd bdev-pokedex
   ```
2. Download dependencies:
   ```sh
   go mod tidy
   ```

### Running the Application

```sh
go build
./pokedex
```

### Running Tests

```sh
go test ./...
```

## Usage

Start the application and use the REPL to enter commands such as:

- `explore <region>`: Explore a Pokémon in the given location
- `catch <pokemon>`: Catch a Pokémon
- `inspect <pokemon>`: View details about a Pokémon that you've caught
- `map`: View a list of 20 map locations
- `mapb`: Go back to the previous 20 locations
- `pokedex`: List all Pokémon you have caught
- `help`: List available commands
- `exit`: Exit the application

## License

MIT License

---

This project is for educational purposes as part of the [Boot.dev](https://www.boot.dev) Go course.
