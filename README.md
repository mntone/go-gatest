# go-gatest

`go-gatest` is a sample repository for validating GitHub Actions + GoReleaser release flows.

## Purpose

This repository is primarily used to verify that GoReleaser can create correct build artifacts and releases across multiple platforms.

## What This Repository Verifies

- Create release artifacts from tag pushes.
- Create snapshot artifacts on pull requests.
- Build a minimal Go CLI app (SQLite-backed task tool) through the same pipeline.

## Task Tool Commands

The CLI is intentionally small and exists mainly as a build target for the workflow.

- `add <description>`: add a task.
- `list`: show all tasks.
- `search <keyword>` (`find` alias): search tasks by keyword.
- `remove <id>`: remove a task by ID.
- `export <path>`: export tasks to a JSON file.
- `import <path>`: import tasks from a JSON file.

## Quick Examples

```bash
go run ./cmd/go-gatest add "check goreleaser snapshot"
go run ./cmd/go-gatest export tasks.json
go run ./cmd/go-gatest import tasks.json
```
