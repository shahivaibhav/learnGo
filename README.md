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

Go modules are used for dependency management. To initialize a module, run:

```bash
go mod init <module-path>
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

## Difference Between `fmt.Println` and `fmt.Printf`

Go's `fmt` package provides multiple output functions. The two most commonly used are `Println` and `Printf`, each serving different purposes.

### `fmt.Println`

`Println` is designed for **simple and readable output**.

Characteristics:

* Automatically adds a newline at the end
* Automatically inserts spaces between arguments
* Does **not** require format specifiers
* Ideal for quick debugging and straightforward logging

Example:

```go
fmt.Println("Name:", name)
fmt.Println(name, age, salary)
```

Typical use cases:

* Debugging
* Printing simple values
* Readable console output where formatting precision is not required

---

### `fmt.Printf`

`Printf` is designed for **precise and controlled formatting**.

Characteristics:

* Does **not** add a newline automatically
* Requires format specifiers (e.g., `%s`, `%d`, `%.2f`, `%c`)
* Does not insert spaces automatically
* Suitable for structured, formatted output

Example:

```go
fmt.Printf("Name: %s\n", name)
fmt.Printf("Salary: %.2f\n", salary)
```

Typical use cases:

* Formatting numbers and strings
* Generating reports
* Logs where output format matters

---

### Key Differences at a Glance

| Aspect             | Println                  | Printf                       |
| ------------------ | ------------------------ | ---------------------------- |
| Newline            | Added automatically      | Must be added manually (`\n`) |
| Spacing            | Automatic                | Manual                       |
| Format specifiers  | Not required             | Required                     |
| Formatting control | Limited                  | High                         |
| Common usage       | Debugging, simple output | Structured, formatted output |

---

### When to Use Which

* Use **`Println`** when simplicity and readability are more important than formatting.
* Use **`Printf`** when you need strict control over how the output is displayed.

Understanding this distinction helps write cleaner, more intentional Go output code.

---



