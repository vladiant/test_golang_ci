# test_golang_ci

[![CI](https://github.com/vladiant/test_golang_ci/actions/workflows/ci.yml/badge.svg)](https://github.com/vladiant/test_golang_ci/actions/workflows/ci.yml)
[![Release](https://github.com/vladiant/test_golang_ci/actions/workflows/release.yml/badge.svg)](https://github.com/vladiant/test_golang_ci/actions/workflows/release.yml)
[![Docker](https://github.com/vladiant/test_golang_ci/actions/workflows/docker.yml/badge.svg)](https://github.com/vladiant/test_golang_ci/actions/workflows/docker.yml)
[![GitLab Pipeline](https://gitlab.com/vladiant/test_golang_ci/badges/main/pipeline.svg)](https://gitlab.com/vladiant/test_golang_ci/-/pipelines)
[![GitLab Coverage](https://gitlab.com/vladiant/test_golang_ci/badges/main/coverage.svg)](https://gitlab.com/vladiant/test_golang_ci/-/jobs)
[![Go Report Card](https://goreportcard.com/badge/github.com/vladiant/test_golang_ci)](https://goreportcard.com/report/github.com/vladiant/test_golang_ci)
[![Coverage Status](https://coveralls.io/repos/github/vladiant/test_golang_ci/badge.svg?branch=main)](https://coveralls.io/github/vladiant/test_golang_ci?branch=main)

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
├── .gitlab-ci.yml                  GitLab CI/CD pipeline
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
| **test** | Runs tests with `-race` and coverage; matrix across Go 1.23 & 1.24; uploads results to Coveralls |
| **build** | Compiles the binary and uploads it as a workflow artifact |
| **security** | Runs `govulncheck` to detect known vulnerabilities |

Concurrent runs for the same branch are automatically cancelled to save CI minutes.

### `release.yml` — runs on `v*` tags

Uses [GoReleaser](https://goreleaser.com) to build cross-platform binaries (Linux, macOS, Windows — amd64 & arm64), generate a changelog, and publish a GitHub Release with checksums.

### `docker.yml` — runs on push to `main` and on `v*` tags

Builds a multi-platform Docker image (`linux/amd64`, `linux/arm64`) and pushes it to the GitHub Container Registry (`ghcr.io`). Uses the GHA cache backend to speed up layer builds. Image tags follow SemVer and are enriched with OCI labels.

### `dependabot.yml`

Automatically opens weekly PRs to update Go module dependencies and GitHub Actions versions.

## GitLab CI/CD pipeline

The `.gitlab-ci.yml` mirrors the GitHub Actions pipeline for teams hosting on GitLab.

| Stage | Job | Description |
|-------|-----|-------------|
| **lint** | `lint` | golangci-lint v2.12.2 |
| **test** | `test:go1.23`, `test:go1.24` | Parallel matrix; coverage report published to MR |
| **build** | `build` | Compiles binary; uploaded as a pipeline artifact |
| **security** | `security` | govulncheck against stdlib and dependencies |
| **release** | `release` | GoReleaser — runs only on `v*` tags |
| **docker** | `docker` | Builds & pushes to GitLab Container Registry; `latest` updated on default branch and tags |

Go module cache is shared across jobs via GitLab's cache keyed on `go.sum`.  
MR pipelines build the Docker image but skip the push (no registry credentials injected).

### Required GitLab CI/CD variables

| Variable | Where to set | Description |
|----------|-------------|-------------|
| `CI_REGISTRY_USER` | Auto-provided | GitLab registry username |
| `CI_REGISTRY_PASSWORD` | Auto-provided | GitLab registry token |
| `CI_REGISTRY` | Auto-provided | Registry host (`registry.gitlab.com`) |
| `GITLAB_TOKEN` | Project → Settings → CI/CD | Personal access token for GoReleaser |

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
