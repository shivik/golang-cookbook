# Go Language Cookbook

This project is an educational codebase that demonstrates the core building blocks of the Go programming language using small, focused, runnable examples. It is inspired by official Go documentation such as the language specification, “A Tour of Go”, and Effective Go, but rewritten in a compact, example‑first form.  

## What this project teaches

Each command under `cmd/` shows a different area of the language:

- `cmd/basics`: basic types, variables, constants, arrays, slices, maps, and structs.
- `cmd/controlflow`: `if`, `switch`, and all forms of `for` loops.
- `cmd/functions`: functions, multiple return values, named results, anonymous functions, closures, and `defer`.
- `cmd/methods_interfaces`: methods (value vs pointer receivers), interfaces, and type embedding.
- `cmd/generics`: type parameters, constraints, and generic helper functions.
- `cmd/concurrency`: goroutines, channels, `select`, buffered vs unbuffered channels, and cancellation with `context`.
- `cmd/modules_packages`: how to organize code into packages within a module and import them.
- `cmd/errors`: Go’s error handling style, sentinel errors, wrapping, and `errors.Is`.

The project aims to make Go’s syntax and mental model easy to grasp by reading and running real code.

## Prerequisites

- Go 1.22 or newer installed and on your `PATH`.  
- A terminal (Linux, macOS, or Windows).  

You can verify your Go installation:


## Getting started

1. **Clone or copy this repository**
2. **Initialize / tidy dependencies**

The module only uses Go’s standard library, but you can still run:
## How to run each topic

Each topic is a separate Go command. To run a topic, change into the project root and use `go run` on the corresponding folder inside `cmd/`.

### Basics

You will see output illustrating:

- How basic types, strings, runes, and bytes behave.
- Different styles of variable and constant declarations.
- Array vs slice behavior (length, capacity) and simple maps.
- Structs and pointers to structs.

### Control flow


You will see:

- `if` and `switch` expressions, including conditionless `switch`.
- Different `for` loop shapes, including `for range`.

### Functions


You will see:

- Multiple return values and named result parameters.
- Error‑returning functions instead of exceptions.
- Closures capturing state.
- How `defer` works and the LIFO order of deferred calls.

### Methods, interfaces, and embedding


You will see:

- Methods on struct types and when to use pointer receivers.
- Interfaces as behavior contracts and dynamic dispatch.
- Type embedding to reuse behavior like a built‑in logger.

### Generics


You will see:

- A reusable `Number` constraint and a generic `SumNumbers` function.
- A generic `Map` helper that converts `[]T` to `[]U`.

### Concurrency


You will see:

- Goroutines running concurrently with the main function.
- Unbuffered and buffered channels, and ranging over a closed channel.
- `select` picking from multiple channel operations.
- Cooperative cancellation with `context.WithTimeout`.

### Modules and packages


You will see:

- Example usage of a small reusable package (`mathutil`) living under `internal/modules_packages/pkgdemo`.

### Errors


You will see:

- Sentinel errors (`ErrNotFound`).
- Checking errors with `errors.Is`.
- Wrapping errors with extra context using `%w`.

## How to extend this project

To cover even more of the language, you can add additional commands under `cmd/` and corresponding packages under `internal/`, for example:

- `cmd/testing` for table‑driven tests and benchmarks.
- `cmd/http` for HTTP servers and clients.
- `cmd/json` for JSON encoding/decoding.
- `cmd/interfaces_advanced` for type assertions and type switches.

Each new topic should:

1. Focus on a single concept.
2. Contain short, readable examples.
3. Print enough output to be self‑explanatory.

This way the repository can grow into a complete Go learning playground.
