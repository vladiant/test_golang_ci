# test_golang_ci

[![CI](https://github.com/vladiant/test_golang_ci/actions/workflows/ci.yml/badge.svg)](https://github.com/vladiant/test_golang_ci/actions/workflows/ci.yml)
[![Release](https://github.com/vladiant/test_golang_ci/actions/workflows/release.yml/badge.svg)](https://github.com/vladiant/test_golang_ci/actions/workflows/release.yml)
[![Docker](https://github.com/vladiant/test_golang_ci/actions/workflows/docker.yml/badge.svg)](https://github.com/vladiant/test_golang_ci/actions/workflows/docker.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/vladiant/test_golang_ci)](https://goreportcard.com/report/github.com/vladiant/test_golang_ci)
[![codecov](https://codecov.io/gh/vladiant/test_golang_ci/branch/main/graph/badge.svg)](https://codecov.io/gh/vladiant/test_golang_ci)

A simple Go project demonstrating CI/CD best practices with GitHub Actions.

## Features

- Basic arithmetic calculator (`add`, `subtract`, `multiply`, `divide`)
- Table-driven unit tests with the race detector enabled
- Full CI/CD pipeline via GitHub Actions

## Project structure

```
.
├── cmd/main.go                     CLI entry point
├── internal/calculator/
│   ├── calculator.go               Business logic
│   └── calculator_test.go          Unit tests
├── .github/
│   ├── dependabot.yml              Automated dependency updates
│   └── workflows/
│       ├── ci.yml                  Main CI pipeline
│       ├── release.yml             Release automation (GoReleaser)
│       └── docker.yml              Docker build & push to GHCR
├── Dockerfile                      Multi-stage distroless image
├── Makefile                        Local development shortcuts
├── .golangci.yml                   Linter configuration
└── .goreleaser.yml                 Cross-platform release configuration
```

## CI/CD pipeline

### `ci.yml` — runs on every push and pull request

| Job | Description |
|-----|-------------|
| **lint** | Runs `golangci-lint` (misspell, revive, gocritic, goimports) |
| **test** | Runs tests with `-race` and coverage; matrix across Go 1.22 & 1.23 |
| **build** | Compiles the binary and uploads it as a workflow artifact |
| **security** | Runs `govulncheck` to detect known vulnerabilities |

Concurrent runs for the same branch are automatically cancelled to save CI minutes.

### `release.yml` — runs on `v*` tags

Uses [GoReleaser](https://goreleaser.com) to build cross-platform binaries (Linux, macOS, Windows — amd64 & arm64), generate a changelog, and publish a GitHub Release with checksums.

### `docker.yml` — runs on push to `main` and on `v*` tags

Builds a multi-platform Docker image (`linux/amd64`, `linux/arm64`) and pushes it to the GitHub Container Registry (`ghcr.io`). Uses the GHA cache backend to speed up layer builds. Image tags follow SemVer and are enriched with OCI labels.

### `dependabot.yml`

Automatically opens weekly PRs to update Go module dependencies and GitHub Actions versions.

## Usage

### Prerequisites

- [Go 1.22+](https://go.dev/dl/)
- [Docker](https://docs.docker.com/get-docker/) (optional, for container builds)

### Run locally

```bash
# Build
make build

# Run
./calc add 3 4        # → 7
./calc divide 10 3    # → 3.333333333

# Test
make test

# Lint
make lint

# Vulnerability check
make vulncheck
```

### Docker

```bash
docker build -t calc .
docker run --rm calc multiply 6 7   # → 42
```

### Release

Push a version tag to trigger GoReleaser:

```bash
git tag v1.0.0
git push origin v1.0.0
```

## License

[MIT](LICENSE)
