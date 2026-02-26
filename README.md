# Lingo

A CLI tool collection for developers. Built with [Cobra](https://github.com/spf13/cobra) and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Installation

```bash
go install github.com/linchen1987/lingo@latest
```

## Usage

### Interactive Mode (TUI)

```bash
lingo
```

Navigate with `↑/k`, `↓/j`, select with `enter`, quit with `q`.

### CLI Commands

#### Datetime

Convert between Unix timestamp and datetime string.

```bash
# Show current timestamp
lingo datetime

# Convert timestamp to datetime
lingo datetime 1772103423

# Convert datetime to timestamp
lingo datetime '2026.02.26 18:57:35 GMT+8'
```

#### Base58

Encode or decode Base58 strings.

```bash
lingo base58 encode hello
lingo base58 decode StV1DL6CwTryKyV
```

#### Base64

Encode or decode Base64 strings.

```bash
lingo base64 encode hello
lingo base64 decode aGVsbG8=
```

## License

MIT
