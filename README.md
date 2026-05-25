# passwordGen

A simple command-line password generator written in Go. Built using [Cobra](https://github.com/spf13/cobra) as a learning project to explore Go CLI development.

## Features

- Generate passwords of any length
- Toggle digits and special characters on or off
- Cross-platform - runs on Linux, macOS, and Windows

## Installation

### Download a pre-built binary

Grab the latest binary for your platform from the [releases page](https://github.com/shobanchiddarth/passwordGen/releases).

Available builds:

| Platform | Architecture | Filename |
|----------|-------------|----------|
| Linux    | amd64        | `passwordGen-linux-amd64-<version>` |
| Linux    | arm64        | `passwordGen-linux-arm64-<version>` |
| macOS    | amd64        | `passwordGen-darwin-amd64-<version>` |
| macOS    | arm64        | `passwordGen-darwin-arm64-<version>` |
| Windows  | amd64        | `passwordGen-windows-amd64-<version>.exe` |

Make the binary executable (Linux/macOS):

```bash
chmod +x passwordGen-linux-amd64-v0.1.0
```

Optionally move it to your PATH:

```bash
sudo mv passwordGen-linux-amd64-v0.1.0 /usr/local/bin/passwordGen
```

### Build from source

Prerequisites: Go 1.22+

```bash
git clone https://github.com/shobanchiddarth/passwordGen.git
cd passwordGen
go build -o passwordGen .
```

To cross-compile for all supported platforms using the build script:

```bash
bash build.bash v0.1.0
```

Binaries will be placed in the `build/` directory, named as `passwordGen-<os>-<arch>-<version>`.

## Usage

```
passwordGen generate [flags]
```

### Flags

| Flag | Shorthand | Default | Description |
|------|-----------|---------|-------------|
| `--length` | `-l` | `12` | Length of the generated password |
| `--digits` | `-d` | `true` | Include digits (0-9) |
| `--special` | `-s` | `true` | Include special characters |

### Examples

Generate a 12-character password (default):

```bash
passwordGen generate
```

Generate a 25-character password with digits and special characters:

```bash
passwordGen generate -l 25 -d -s
```

Generate a 16-character alphabetic-only password:

```bash
passwordGen generate -l 16 -d=false -s=false
```
