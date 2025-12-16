# learnGo

This repository is intended for learning and practicing Go (Golang) fundamentals.

## The `main` Package

In Go, the `main` package is the entry point of an application.

### Types of Go Programs

There are two primary types of programs in Go:

1. **Executable Programs**

   * Programs that can be built and executed directly from the terminal.
   * Must define a `main` package and a `main` function.

2. **Library Programs**

   * Reusable packages that expose functionality to other Go programs.
   * Do not have a `main` function and cannot be executed directly.

### Requirements for an Executable Program

* The package name must be `main`.
* A `main()` function must be defined.
* The `main()` function:

  * Does not take any parameters.
  * Does not return any value.

Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go")
}
```

## Initializing a Go Module

Go modules are used for dependency management.

To initialize a module, run:

```bash
go mod init <module_name>
```

This command:

* Creates a `go.mod` file at the root of the project.
* Defines the module path.
* Ensures all imported packages in the project are resolved relative to this module.

Example:

```bash
go mod init github.com/username/learnGo
```

Once initialized, any packages you create or import will use this module path as their base.

---

This repository will progressively cover Go basics, standard library usage, and best practices.
