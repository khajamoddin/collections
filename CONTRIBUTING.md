# Contributing Guide

Thanks for your interest in improving `github.com/khajamoddin/collections`! This project aims to be a reliable, idiomatic Go collections library. Contributions of all kinds are welcome: bug fixes, features, docs, and tests.

## Getting Started

- Go 1.22+ is required.
- Clone the repo and install dependencies (standard library only).
- Keep code ASCII and run `gofmt -w` on touched files.

## Development Workflow

1. Create a branch.
2. Write code and tests. Aim for meaningful coverage and add regression tests for bugs.
3. Run `go test ./...` (set `GOCACHE=$(pwd)/.gocache` if your environment limits default cache locations).
4. Run `gofmt -w` and `go vet ./...` before sending a PR.
5. Update docs (`README.md`, `docs/usage.md`, `docs/value.md`, etc.) when APIs change.
6. If your change affects performance, add/adjust benchmarks.
7. Open a PR using the provided template and fill out the checklist.

## Coding Standards

- Keep APIs minimal and zero-value safe where practical.
- Prefer total functions (avoid panics) and document edge cases.
- Maintain ordered iteration guarantees for `OrderedMap`.
- Keep deques O(1) amortized for front/back operations.
- Avoid breaking changes; if unavoidable, call them out clearly and update the changelog.

## Communication

- Use GitHub Issues for bugs and feature requests.
- Security concerns: see `SECURITY.md`.
- Be respectful and follow the `CODE_OF_CONDUCT.md`.
