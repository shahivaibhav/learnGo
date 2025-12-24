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
   age := 19
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

## Taking User Input in Go (`fmt.Scan`, `fmt.Scanln`, `fmt.Scanf`)

Go provides multiple functions in the `fmt` package to read user input from standard input. Each behaves differently and is suitable for specific scenarios.

### `fmt.Scan`

`Scan` reads space-separated values from standard input.

Characteristics:

* Stops reading at whitespace
* Can read multiple values in one call
* Requires pointers to variables
* Ignores newlines

Example:

```go
var firstName, lastName string
fmt.Scan(&firstName, &lastName)
```

Best used when:

* Input values are separated by spaces
* Exact formatting is not required

---

### `fmt.Scanln`

`Scanln` reads input until a newline is encountered.

Important behavior:

* Stops scanning at the first space
* Treats spaces as separators
* Does **not** read an entire line into a single string

```go
var name string
fmt.Scanln(&name)
```

This is why input like:

```
Vaibhav Shahi
```

Results in:

```
name = "Vaibhav"
```

The remaining input (`Shahi`) stays in the buffer and can interfere with subsequent scans.

---

### `fmt.Scanf`

`Scanf` reads input based on a format string.

Characteristics:

* Requires format specifiers
* Input must strictly match the format
* Fails silently if the format does not match

Example:

```go
var age int
var city string
fmt.Scanf("%d %s", &age, &city)
```

If input does not match the expected format, variables remain unchanged (zero values).

This explains output such as:

```text
You are 0 years old and live in .
```

---

### Common Pitfalls

1. **Leftover input in buffer**
   * Mixing `Scan`, `Scanln`, and `Scanf` can leave unread input
   * This causes unexpected behavior in subsequent reads

2. **Spaces in input**
   * None of these functions read full lines containing spaces into a single string

---

### Recommended Approach for Full-Line Input

For reading an entire line (including spaces), use `bufio.Reader`:

```go
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n')
```

This is the idiomatic and reliable way to handle user input in real-world Go programs.

---

### Summary

| Function | Reads spaces          | Format-based | Common use case               |
| -------- | --------------------- | ------------ | ----------------------------- |
| Scan     | No                    | No           | Simple, space-separated input |
| Scanln   | No                    | No           | Single token per line         |
| Scanf    | Depends on format     | Yes          | Structured input              |

For production-grade input handling, prefer `bufio.Reader` over `fmt.Scan*` functions.

---

## User Input Using `bufio.Reader`

For real-world Go applications, the recommended way to take user input—especially full-line input containing spaces—is by using the `bufio` package.

### Why `bufio.Reader` Is Preferred

`bufio.Reader` provides buffered, line-oriented input handling, which solves many limitations of `fmt.Scan`, `Scanln`, and `Scanf`.

Key advantages:

* Reads complete lines, including spaces
* Avoids leftover input buffer issues
* More efficient for large or repeated input
* Works well for both console input and file input

---

### Basic Usage

```go
reader := bufio.NewReader(os.Stdin)
input, err := reader.ReadString('\n')
```

Explanation:

* `bufio.NewReader` wraps `os.Stdin` with a buffer
* `ReadString('\n')` reads input until a newline character is encountered
* Returns the input string and an error (usually `nil` for valid input)

---

### Example: Reading Full-Line Input

```go
reader := bufio.NewReader(os.Stdin)
fmt.Println("Enter your address:")
address, _ := reader.ReadString('\n')
fmt.Println("Your address is:", address)
```

This correctly reads input such as:

```
221B Baker Street, London
```

Without truncation or buffer side effects.

---

### Commonly Used `bufio.Reader` Methods

| Method              | Description                                            |
| ------------------- | ------------------------------------------------------ |
| `ReadString(delim)` | Reads input until the delimiter is found               |
| `ReadBytes(delim)`  | Similar to `ReadString`, returns bytes                 |
| `ReadLine()`        | Low-level line reading (not recommended for beginners) |

---

### When to Use Each Input Method

| Use case                           | Recommended approach |
| ---------------------------------- | -------------------- |
| Simple numeric or token input      | `fmt.Scan`           |
| Strictly formatted input           | `fmt.Scanf`          |
| Full-line input (names, addresses) | `bufio.Reader`       |
| Production / real applications     | `bufio.Reader`       |

---

This approach is idiomatic Go and should be your default choice whenever input correctness and reliability matter.

## Important `bufio.Reader` Methods

The `bufio.Reader` type provides buffered input operations and is commonly used for reading data from standard input, files, and network connections. It improves performance by reducing the number of underlying read system calls and provides convenient methods for line-based and byte-level input.

---

### `ReadString(delim byte)`

Reads input until the specified delimiter is encountered.

```go
input, err := reader.ReadString('\n')
```

Characteristics and behavior:

* Returns a string
* Includes the delimiter in the returned value
* Blocks until the delimiter is found or an error occurs
* Suitable for line-oriented input

Common use cases:

* Reading full lines from standard input
* Processing text files line by line
* Interactive CLI programs

---

### `ReadBytes(delim byte)`

Reads input until the specified delimiter and returns the data as a byte slice.

```go
data, err := reader.ReadBytes('\n')
```

Characteristics and behavior:

* Returns `[]byte`
* Includes the delimiter
* Avoids string allocation
* More efficient for byte-level processing

Common use cases:

* Handling binary or streamed data
* Network protocol parsing
* Performance-sensitive applications

---

### `ReadLine()`

Reads a single line from the buffer.

```go
line, isPrefix, err := reader.ReadLine()
```

Characteristics and behavior:

* Returns the line without the newline character
* `isPrefix` is true if the line is longer than the internal buffer
* Can return partial lines that must be manually combined
* Provides low-level control over input reading

Common use cases:

* Building custom parsers
* Reading very large lines with controlled buffering
* Situations where memory usage must be tightly controlled

**Important note:** `ReadLine` is not recommended for beginners because improper handling of `isPrefix` can lead to truncated or corrupted input.

---

### `Peek(n int)`

Returns the next `n` bytes without advancing the reader.

```go
data, err := reader.Peek(5)
```

Characteristics and behavior:

* Does not consume the data
* Fails if fewer than `n` bytes are available
* Allows inspection of upcoming input

Common use cases:

* Checking prefixes or headers
* Protocol detection
* Conditional parsing logic

---

### `Discard(n int)`

Discards the next `n` bytes from the buffer.

```go
reader.Discard(10)
```

Characteristics and behavior:

* Skips unwanted data
* Advances the reader position
* Efficient way to ignore input

Common use cases:

* Ignoring headers or metadata
* Skipping malformed input
* Efficiently bypassing data

---

### `Reset(r io.Reader)`

Resets the `bufio.Reader` to read from a new underlying `io.Reader`.

```go
reader.Reset(os.Stdin)
```

Characteristics and behavior:

* Reuses the existing buffer
* Avoids reallocating memory
* Allows changing the input source dynamically

Common use cases:

* Reusing readers in long-running applications
* Switching between different input streams
* Optimizing memory usage

---

### Summary of `bufio.Reader` Methods

| Method       | Return Type | Includes Delimiter | Typical Use Case                      |
| ------------ | ----------- | ------------------ | ------------------------------------- |
| ReadString   | string      | Yes                | Line-based text input                 |
| ReadBytes    | []byte      | Yes                | Byte-level and binary processing      |
| ReadLine     | []byte      | No                 | Low-level controlled input reading    |
| Peek         | []byte      | No                 | Lookahead without consuming input     |
| Discard      | int         | No                 | Skipping unwanted data                |
| Reset        | —           | —                  | Reusing reader with new input source  |

---

### Best Practice

For most Go applications:

* Prefer `ReadString('\n')` for clarity and correctness
* Use `ReadBytes` when working with raw or binary data
* Avoid `ReadLine` unless you fully understand its behavior and limitations

Mastering these `bufio.Reader` methods is essential for writing reliable and production-quality Go programs.

# Go Functions Guide

A comprehensive guide to understanding and using functions in Go, covering various syntax patterns and best practices.

## Overview

This repository demonstrates different ways to declare and use functions in Go, from basic syntax to advanced features like multiple return values and named returns.

## Function Syntax

```go
func functionName(parameters) returnType {
    function body
    return statement
}
```

## Examples Covered

### 1. Function with No Parameters and No Return Type
```go
func addTwoNumbers() {
    sum := 5 + 10
    println("Sum is:", sum)
}
```
- Simplest form of function
- Performs an operation without taking input or returning output
- Useful for self-contained operations

### 2. Function with Parameters and No Return Type
```go
func addNumbers(a int, b int) {
    res := a + b
    println("Result is:", res)
}
```
- Accepts input parameters
- Performs operation and outputs result directly
- No return value needed

### 3. Function with Parameters and Return Type
```go
func multiplyNumbers(a int, b int) int {
    return a * b
}
```
- Accepts parameters and returns a value
- Most common function pattern
- Return type must be specified

### 4. Cleaner Syntax for Same-Type Parameters
```go
func betterMultiplyNumbers(a, b int) int {
    return a * b
}
```
- When multiple parameters share the same type, declare the type once
- More concise and readable
- Recommended Go style

### 5. Multiple Return Values
```go
func swapValues(a string, b string) (string, string) {
    return b, a
}
```
- Go supports returning multiple values
- Useful for returning results along with error status
- Common pattern in Go standard library

### 6. Named Return Values
```go
func divideNumbers(a, b int) (quotient int, remainder int) {
    quotient = a / b
    remainder = a % b
    return
}
```
- Return values can be named in the function signature
- Acts as documentation
- Naked `return` statement returns all named values
- Useful for complex functions with multiple returns

## Running the Demo

```go
func LearnFunctions() {
    println("----- Function Demo -----")
    addTwoNumbers()
    addNumbers(20, 30)
    result := multiplyNumbers(5, 6)
    println("Multiplication Result is:", result)
    betterResult := betterMultiplyNumbers(7, 8)
    println("Better Multiplication Result is:", betterResult)
    x, y := swapValues("Hello", "World")
    println("After Swapping: x =", x, ", y =", y)
    q, r := divideNumbers(17, 5)
    println("Quotient:", q, ", Remainder:", r)
}
```

### Expected Output
```
----- Function Demo -----
Sum is: 15
Result is: 50
Multiplication Result is: 30
Better Multiplication Result is: 56
After Swapping: x = World , y = Hello
Quotient: 3 , Remainder: 2
```

## Key Takeaways

- **Simplicity**: Go functions are straightforward and easy to understand
- **Multiple Returns**: Leverage multiple return values for cleaner error handling
- **Named Returns**: Use named returns for better documentation in complex functions
- **Type Grouping**: Group parameters of the same type for cleaner syntax
- **Flexibility**: Functions can have zero to multiple parameters and return values

## Best Practices

1. Use named return values sparingly - they're best for complex functions where they improve clarity
2. Group parameters of the same type using the shorter syntax
3. Return errors as the last return value (idiomatic Go pattern)
4. Keep functions focused on a single responsibility
5. Use descriptive function names that indicate what the function does

## Learning Path

Start with simple functions and gradually move to more complex patterns:
1. No parameters, no return → Basic operations
2. With parameters, no return → Procedures
3. With parameters and return → Pure functions
4. Multiple returns → Error handling patterns
5. Named returns → Complex logic documentation

# Go Underscore Identifier Guide

A comprehensive guide to understanding and using the blank identifier (`_`) in Go, covering various use cases and best practices.

## Overview

In Go, an underscore (`_`) is known as the **blank identifier**. It is a special identifier used to ignore values that are not needed in your code. The blank identifier can be assigned any value of any type, and that value is discarded.

## When to Use the Blank Identifier

The blank identifier is commonly used in the following scenarios:

1. **Ignoring return values from functions**
2. **Ignoring loop indices or values**
3. **Importing packages for side effects only**
4. **Interface implementation checks**
5. **Unused variables during development**

## Examples Covered

### Example 1: Ignoring Return Values from a Function

```go
func divideTwoNumbers(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

**Ignoring the error (not recommended for production):**
```go
a, _ := divideTwoNumbers(10, 2)
fmt.Println("Result of division is:", a)
```

**Ignoring the result:**
```go
_, err := divideTwoNumbers(10, 0)
if err != nil {
    fmt.Println("Error occurred:", err)
}
```

**Proper error handling (recommended):**
```go
c, err := divideTwoNumbers(10, 0)
if err != nil {
    fmt.Println("Error occurred:", err)
} else {
    fmt.Println("Result of division is:", c)
}
```

### Example 2: Ignoring Index in a Loop

```go
func ignoreIndexInLoop() {
    fmt.Println("----- Ignoring Index in Loop Demo -----")
    numbers := []int{10, 20, 30, 40, 50}
    sum := 0

    for _, value := range numbers {
        sum += value
    }

    fmt.Println("Sum of numbers is:", sum)
}
```

**When you only need the index:**
```go
for index, _ := range numbers {
    fmt.Println("Index:", index)
}
```

**Or more idiomatically:**
```go
for index := range numbers {
    fmt.Println("Index:", index)
}
```

## Running the Demo

```go
func UnderscoreIdentifierDemo() {
    fmt.Println("----- Underscore Identifier Demo -----")
    
    a, _ := divideTwoNumbers(10, 2)
    fmt.Println("Result of division is:", a)

    b, _ := divideTwoNumbers(10, 0)
    fmt.Println("Result of division is:", b)

    c, err := divideTwoNumbers(10, 0)
    if err != nil {
        fmt.Println("Error occurred:", err)
    } else {
        fmt.Println("Result of division is:", c)
    }

    ignoreIndexInLoop()
}
```

### Expected Output

```
----- Underscore Identifier Demo -----
Result of division is: 5
Result of division is: 0
Error occurred: division by zero
----- Ignoring Index in Loop Demo -----
Sum of numbers is: 150
```

## Additional Use Cases

### 3. Import for Side Effects Only

```go
import (
    _ "database/sql/driver" // Import for driver registration
)
```

### 4. Interface Implementation Check

```go
var _ InterfaceName = (*StructName)(nil)
```

This ensures at compile time that `StructName` implements `InterfaceName`.

### 5. Unused Variables During Development

```go
func example() {
    x := 10
    _ = x // Temporarily ignore x to avoid compilation errors
    
    // TODO: Use x later
}
```

## Best Practices

### ✅ DO:

1. **Use `_` when you genuinely don't need a value**
   ```go
   for _, value := range items {
       process(value)
   }
   ```

2. **Use it to ction()  make your intent clear**
   ```go
   _, err := someFun// Clear that we only care about the error
   ```

3. **Use it for interface checks**
   ```go
   var _ io.Reader = (*MyReader)(nil)
   ```

### ❌ DON'T:

1. **Ignore errors in production code**
   ```go
   // BAD - ignoring potential errors
   result, _ := riskyOperation()
   
   // GOOD - handle errors properly
   result, err := riskyOperation()
   if err != nil {
       log.Fatal(err)
   }
   ```

2. **Overuse it when a variable name would be clearer**
   ```go
   // Less clear
   for _, _ := range items {
       // ...
   }
   
   // Better
   for range items {
       // ...
   }
   ```

3. **Use it to silence the compiler for genuinely unused variables**
   - If a variable is unused, consider removing it instead

## Common Patterns

### Pattern 1: Multiple Return Values

```go
// Only need the first value
value, _ := functionReturningTwo()

// Only need the second value
_, status := functionReturningTwo()

// Only need middle values
_, middle, _ := functionReturningThree()
```

### Pattern 2: Range Loops

```go
// Only values needed
for _, value := range slice {
    fmt.Println(value)
}

// Only indices needed
for index := range slice {  // Cleaner than: for index, _ := range slice
    fmt.Println(index)
}

// Neither needed (just counting iterations)
for range slice {
    count++
}
```

### Pattern 3: Map Operations

```go
// Check if key exists, ignore value
if _, exists := myMap[key]; exists {
    fmt.Println("Key exists")
}

// Get value, ignore existence check
value := myMap[key]  // Returns zero value if not exists
```

## Key Takeaways

- The blank identifier `_` is a write-only variable that discards any value assigned to it
- It helps satisfy Go's requirement that all declared variables must be used
- **Never ignore errors in production code** - always handle them appropriately
- Use `_` to make your code's intent clear and more readable
- When you only need the index in a range loop, omit the blank identifier: `for i := range items`

## Usage

To use this code in your project:

1. Import the package:
```go
import "yourmodule/underscoreidentifier"
```

2. Call the demo function:
```go
func main() {
    underscoreidentifier.UnderscoreIdentifierDemo()
}
```

## Warning

While the blank identifier is useful, **be especially careful when ignoring error values**. In the example above, ignoring the error from `divideTwoNumbers(10, 0)` results in using a zero value, which might not be the intended behavior. Always handle errors appropriately in production code.
