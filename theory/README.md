# Go theory notes

These short notes introduce the ideas behind the repository examples. Follow the sequence below: read a note, open its linked Go file, inspect the comments and example function, then try the corresponding call from `main.go` and run `go run .` at the repository root. The examples share one `package main`; only one example call is enabled at a time.

## Learning path

1. [Getting started, packages, and basic types](01-foundations.md) — `hellow_world.go`, `main.go`, `go.mod`, `basicTypes.go`
2. [Control flow and collections](02-control-flow-and-collections.md) — `loops_case.go`, `arrays.go`, `slices.go`, `maps.go`
3. [Functions, strings, and structs](03-functions-strings-structs.md) — `functions.go`, `strings_rune.go`, `struct_method.go`
4. [Concurrency](04-concurrency.md) — `go_routine.go`, `channels.go`, `select_timeout.go`, `worker_pools.go`, `rate_limitter.go`
5. [Errors and program flow](05-errors-and-program-flow.md) — `errors.go`, `panic_defer_recover.go`
6. [Standard library essentials](06-standard-library.md) — `random.go`, `time.go`, `regularexp.go`, `json.go`, `logging.go`
7. [Processes and the web](07-processes-and-web.md) — `commandLineArgs.go`, `environmentVariable.go`, `urlParsing.go`, `http_call_api.go`, `go-create_server.go`

## Suggested practice

Run one example at a time, predict what it will print or where it will wait, then run it and compare. Change a value or add a small case, and revisit the note when the behavior surprises you. Some files intentionally show a blocking mistake or use external network access; read the file before enabling it. `go-create_server.go` starts a local server on port 8080 and keeps running until stopped.
