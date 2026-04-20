# sparke

`sparke` is a personal CLI tool to improve productivity when starting new projects.

It scaffolds simple project structures for:
- Rust
- Python
- Go

## Requirements

This tool runs external commands, so you need these installed:
- Go (to build and run `sparke`)
- `git` (to initialize generated repositories)
- `faz` (to initialize task tracking in generated projects)
- [Cargo](https://github.com/rust-lang/cargo) (for Rust project creation)
- [uv](https://github.com/astral-sh/uv) (for Python project creation)
- [just](https://github.com/casey/just) (optional, to use the generated `justfile` recipes)

## Quick start

```bash
# Install directly from GitHub:
go install github.com/rpcarvs/sparke@latest

# Or clone the repo and install locally:
go install .
# If `sparke` is not found, add GOPATH/bin to PATH (Bash example):
grep -q '$(go env GOPATH)/bin' ~/.bashrc || echo 'export PATH="$PATH:$(go env GOPATH)/bin"' >> ~/.bashrc
source ~/.bashrc
```

## Usage

Run `sparke` with no subcommand to enter the interactive flow.

```bash
sparke
```

The interactive flow will:
- let you choose Go, Rust, or Python
- ask for Rust and Python project types when applicable
- ask for the target directory, where an empty value means the current directory
- ask separately for the Go module name when Go is selected

You can still use the explicit CLI commands when you want a non-interactive flow.

Useful built-in CLI helpers:

```bash
sparke --version
sparke completion zsh
```

### Go project

To scaffold a Go project with a minimalist dir structure:

```bash
sparke go my-go-app
```

### Rust project

To scaffold a Rust project (binary or library) following `cargo init` standards:

```bash
sparke rust my-rust-app
sparke rust my-rust-lib --lib
```


### Python project

To scaffold a Python project (app, library or package) following `uv init` standards:

```bash
sparke python my-python-app
sparke python my-python-lib --lib
sparke python my-python-package --package
```
