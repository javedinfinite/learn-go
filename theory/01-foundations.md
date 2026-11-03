# 1. Getting Started, Packages, and Basic Types

In Go, code is organized into **packages**. A package is simply a way to group related Go files.

To create a runnable Go application:

- `package main` tells Go that this package is an executable program.
- `func main()` is the function where program execution starts.
- `go run .` runs the Go program in the current directory.
- `go.mod` contains basic project information, such as the module name and Go version.

### Example files

- [`hellow_world.go`](../hellow_world.go) — A simple example that prints text.
- [`main.go`](../main.go) — Contains the main function and lets you run different examples.
- [`go.mod`](../go.mod) — Defines the Go module and version.
- [`basicTypes.go`](../basicTypes.go) — Demonstrates variables, data types, default values, and formatted output.

### Variables and basic types

Go is a **statically typed language**, meaning every variable has a specific type.

Go can automatically determine the type from the assigned value:

    n := 3        // int
    name := "Go"  // string
    active := true // bool

You can also declare the type explicitly:

    var age int = 25

Or declare a variable without assigning a value:

    var count int

Go automatically assigns a **zero value**:

- Numbers → `0`
- Boolean → `false`
- String → `""` (empty string)

### Strings and runes

- `string` — Stores text as UTF-8 encoded bytes.
- `rune` — Represents a Unicode code point (internally an `int32`).

### Advanced note

The `unsafe` package allows low-level memory operations. Converting a pointer to `uintptr` treats its address as an integer, but it should not be used like an ordinary type conversion because pointer safety rules are different.

**Summary:** Go organizes code into packages, starts executable programs through `main()`, supports automatic type inference, and provides default values for variables.