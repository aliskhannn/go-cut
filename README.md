# gocut

`gocut` is a command-line utility to extract specific columns from text, similar to the Unix `cut` command.

It supports:

* Selecting fields (`-f "1,3-5"`)
* Custom delimiter (`-d ":"`)
* `-s` flag to output only lines containing the delimiter

## Installation

```bash
git clone https://github.com/aliskhannn/go-cut.git
cd go-cut
make build
```

The binary will be available at `bin/gocut`.

## Project Structure

```
go-cut/
├── cmd/
│   └── gocut/             # main.go for building the binary
├── internal/
│   └── cut/               # Cut utility implementation
│       ├── flags.go       # CLI flags parsing
│       ├── cut.go         # Core logic
│       └── cut_test.go    # Unit tests
├── integration/           # integration tests
├── testdata/              # test files for integration tests
├── Makefile               # build, test, lint, clean commands
└── README.md
```

## Usage

### Basic

```bash
bin/gocut -f 1,3 -d ":" testdata/basic.txt
# Output:
# a:c
# e:g
# i:k
```

### Only lines with delimiter

```bash
bin/gocut -f 2 -d ":" -s testdata/separated_only.txt
# Output:
# y
# q
```

### Reading from stdin

```bash
echo "1,2,3\n4,5,6" | bin/gocut -f 2 -d ","
# Output:
# 2
# 5
```

### Using comma-separated file

```bash
bin/gocut -f 1,3 -d "," testdata/comma.csv
# Output:
# 1,3
# 4,6
```

## Makefile Commands

* `make build` — build the binary
* `make test` — run unit and integration tests
* `make lint` — check code with golangci-lint and go vet
* `make clean` — remove binaries and temporary files