---
layout: default
title: Continuous Integration
nav_order: 6
---

# Continuous Integration

Workflow: `.github/workflows/go.yml`

- Triggers: `push`, `pull_request`
- Steps:
  - `actions/checkout@v3`
  - `actions/setup-go@v4` with `go-version: 1.22`
  - `go mod tidy`
  - `go test ./... -v -cover`
  - `go build ./...`

## Badges

Once connected to GitHub, add badges for build status and coverage to `README.md`.
