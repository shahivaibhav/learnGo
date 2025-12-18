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

## Variables and Constants in Go

This section explains how variables and constants work in Go, along with common declaration patterns and scoping rules.

### `var` vs `const`

* **`var`** is used to declare variables whose values can be changed during program execution.
* **`const`** is used to declare constants whose values are fixed at compile time and cannot be modified later.

Example:

```go
var age int = 25
const country string = "India"
```

### Ways to Declare Variables

Go supports multiple ways to declare variables:

1. **Explicit type declaration**

```go
var name string = "Vaibhav"
```

2. **Type inference** (compiler infers the type from the value)

```go
var isDeveloper = true
```

3. **Short variable declaration (`:=`)**

```go
take := 19
```

> Note: Short variable declaration (`:=`) is **only allowed inside functions**. It is not permitted at the package level.

### Declaring Multiple Variables

Multiple variables can be declared together using a `var` block:

```go
var (
    city        string = "Gurgaon"
    countryCode int    = 91
)
```

Similarly, constants can be grouped using a `const` block:

```go
const (
    pi       = 3.14
    language = "Go"
)
```

### Package-Level Declarations

* Variables and constants declared outside functions are **package-level**.
* Package-level declarations must use `var` or `const`.
* `:=` cannot be used at the package level because it relies on runtime context, which is only available inside functions.

### Printing Variables

The `fmt` package is commonly used to print variables:

```go
fmt.Println("Name:", name)
```

### Exported vs Unexported Identifiers

In Go, visibility is controlled by naming:

* **Exported** identifiers start with an uppercase letter and are accessible outside the package.
* **Unexported** identifiers start with a lowercase letter and are accessible only within the same package.

Example:

```go
func ExportedFunction() {
    fmt.Println("This is an exported function.")
}
```

---

This module demonstrates best practices for declaring variables and constants in Go and highlights important rules around scope, immutability, visibility, and compile-time behavior.

## Additional Concepts for Variables and Constants

### Zero Values

In Go, variables that are declared but not explicitly initialized are automatically assigned **zero values**:

```go
var i int       // 0
var s string    // ""
var b bool      // false
```

This eliminates the risk of uninitialized variables and is a key language design feature.

### Compile-Time Nature of Constants

Constants in Go are evaluated at **compile time**.

```go
const x = 10
const y = len("go") // valid
```

A constant **cannot** be assigned a value returned from a function:

```go
// invalid
const z = getValue()
```

### Untyped Constants

By default, Go constants are **untyped**, allowing flexible assignment:

```go
const n = 10

var a int = n
var b float64 = n
```

The constant adapts its type based on the context in which it is used.

### Variable Shadowing

A variable declared in an inner scope can **shadow** a variable with the same name from an outer scope:

```go
x := 10

if true {
    x := 20 // shadows outer x
}
```

Shadowing is legal but should be used carefully to avoid logical errors.

### Short Variable Declaration Rules

The short declaration operator (`:=`) has a redeclaration rule:

```go
a := 10
a, b := 20, 30 // valid (at least one new variable)
```

However:

```go
// invalid
a := 10
a := 20
```

At least one variable on the left-hand side must be new.

### Using `iota` with Constants

`iota` is used to create incrementing constants:

```go
const (
    A = iota // 0
    B        // 1
    C        // 2
)
```

This pattern is commonly used to define enums in Go.

### When to Use `var` Instead of `:=`

Use `var` when:

* You want the zero value
* You need explicit typing
* You are declaring variables at the package level

Use `:=` when:

* Declaring and initializing variables inside functions
* Writing concise and idiomatic Go code

---

## Difference Between `fmt.Println` and `fmt.Printf`

Go’s `fmt` package provides multiple output functions. The two most commonly used are `Println` and `Printf`, each serving different purposes.

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

| Aspect             | Println                  | Printf                        |
| ------------------ | ------------------------ | ----------------------------- |
| Newline            | Added automatically      | Must be added manually (`\n`) |
| Spacing            | Automatic                | Manual                        |
| Format specifiers  | Not required             | Required                      |
| Formatting control | Limited                  | High                          |
| Common usage       | Debugging, simple output | Structured, formatted output  |

---

### When to Use Which

* Use **`Println`** when simplicity and readability are more important than formatting.
* Use **`Printf`** when you need strict control over how the output is displayed.

Understanding this distinction helps write cleaner, more intentional Go output code.
