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
