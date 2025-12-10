# Advent of Code 2025 🎄

My solutions for [Advent of Code 2025](https://adventofcode.com/2025) written in Go.

## Project Structure

```
.
├── day01/              # Day 1 solution
│   ├── main.go         # Solution code
│   ├── input.txt       # Puzzle input
│   ├── example.txt     # Example input from problem statement
│   └── README.md       # Day-specific notes
├── day02/              # Day 2 solution
│   └── ...
├── ...
├── day25/              # Day 25 solution
│   └── ...
├── utils/              # Common utility functions
│   └── utils.go        # File reading, parsing helpers, etc.
├── Makefile            # Build and run commands
├── go.mod              # Go module definition
└── README.md           # This file
```

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

```bash
git clone https://github.com/UglyBoySadFace/advent-of-code-2025.git
cd advent-of-code-2025
```

## Usage

### Running Solutions

Run a specific day's solution:

```bash
make run DAY=01
```

Or navigate to a day's directory and run directly:

```bash
cd day01
go run main.go
```

### Running Tests

Run all tests:

```bash
make test
```

Run tests for a specific day:

```bash
make test DAY=01
```

### Other Commands

Format all code:

```bash
make fmt
```

Clean build artifacts:

```bash
make clean
```

View all available commands:

```bash
make help
```

## Utilities

The `utils` package provides common helper functions:

- `ReadLines(filename)` - Read file as slice of lines
- `ReadFile(filename)` - Read entire file as string
- `MustReadLines(filename)` - Read lines or panic
- `MustReadFile(filename)` - Read file or panic
- `ToIntSlice([]string)` - Convert strings to integers
- `MustToIntSlice([]string)` - Convert to ints or panic
- `Abs(int)` - Absolute value
- `Min(a, b int)` - Minimum of two integers
- `Max(a, b int)` - Maximum of two integers
- `Sum([]int)` - Sum of slice

## Workflow

For each day:

1. Read the problem at https://adventofcode.com/2025/day/X
2. Add your puzzle input to `dayXX/input.txt`
3. Add example input to `dayXX/example.txt` (optional, for testing)
4. Implement solutions in `dayXX/main.go`
5. Run with `make run DAY=XX`

## Progress

- [ ] Day 01
- [ ] Day 02
- [ ] Day 03
- [ ] Day 04
- [ ] Day 05
- [ ] Day 06
- [ ] Day 07
- [ ] Day 08
- [ ] Day 09
- [ ] Day 10
- [ ] Day 11
- [ ] Day 12
- [ ] Day 13
- [ ] Day 14
- [ ] Day 15
- [ ] Day 16
- [ ] Day 17
- [ ] Day 18
- [ ] Day 19
- [ ] Day 20
- [ ] Day 21
- [ ] Day 22
- [ ] Day 23
- [ ] Day 24
- [ ] Day 25

## License

This project is open source and available under the MIT License.