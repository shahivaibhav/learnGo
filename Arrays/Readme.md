 # Go Arrays Guide

A comprehensive guide to understanding and working with arrays in Go, covering declaration methods, initialization, and important array concepts.

## Overview

Arrays in Go are fixed-size sequences of elements of the same type. Unlike slices, arrays have a fixed length that is part of their type, making `[5]int` and `[10]int` two completely different types.

## Array Declaration Methods

### Method 1: Declare an Array with Fixed Size

```go
var arr1 [5]int   // Integer array of size 5
fmt.Println(arr1) // Output: [0 0 0 0 0]

var arr2 [3]string       // String array of size 3
fmt.Printf("%q\n", arr2) // Output: ["" "" ""]
```

**Key Points:**
- Must specify both size and type
- Default values: `0` for numeric types, `""` for strings, `false` for booleans, `nil` for pointers

**Assigning Values:**
```go
arr1[0] = 10
arr1[1] = 20
arr1[2] = 30

arr2[0] = "Hello"
arr2[1] = "World"

fmt.Println(arr1)        // Output: [10 20 30 0 0]
fmt.Printf("%q\n", arr2) // Output: ["Hello" "World" ""]
```

### Method 2: Declare and Initialize in One Line

```go
arr3 := [6]int{1, 2, 3, 4, 5, 6}
fmt.Println(arr3) // Output: [1 2 3 4 5 6]
```

**Key Points:**
- Use curly braces `{}` for initialization
- Short declaration syntax with `:=`
- Must provide size explicitly

### Method 3: Let Go Determine the Size (Ellipsis)

```go
arr4 := [...]string{"Go", "is", "awesome"}
fmt.Printf("%q\n", arr4) // Output: ["Go" "is" "awesome"]
```

**Key Points:**
- Use `...` instead of a fixed size
- Go automatically counts the elements
- Still creates a fixed-size array at compile time
- Type becomes `[3]string` in this example

### Method 4: Multidimensional Arrays (2D Arrays)

```go
var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
fmt.Println(arr) // Output: [[1 2 3] [4 5 6]]
```

**Key Points:**
- Format: `[rows][columns]type`
- Can have multiple dimensions
- Each dimension must have a fixed size

**Accessing 2D Array Elements:**
```go
fmt.Println(arr[0][0]) // Output: 1 (first row, first column)
fmt.Println(arr[1][2]) // Output: 6 (second row, third column)
```

## Important Array Concepts

### 1. Arrays are Value Types

When you assign an array to another variable, Go creates a **complete copy** of the array.

```go
arr5 := arr3      // arr5 is a copy of arr3
arr5[0] = 100

fmt.Println("arr3:", arr3) // Output: arr3: [1 2 3 4 5 6]
fmt.Println("arr5:", arr5) // Output: arr5: [100 2 3 4 5 6]
```

**Implications:**
- Modifying the copy doesn't affect the original
- Passing arrays to functions creates copies (can be expensive for large arrays)
- Use pointers or slices if you need to modify the original

### 2. Array Length

Use the `len()` function to get the length of an array:

```go
fmt.Println("Length of arr4:", len(arr4)) // Output: Length of arr4: 3
```

**Important:** The length is fixed and part of the array's type.

### 3. Contiguous Memory Storage

Arrays are stored in contiguous memory locations, which allows for:
- Efficient access using indices
- Cache-friendly operations
- Predictable memory layout

```
Memory Layout of [5]int:
[0][1][2][3][4]
 ↓  ↓  ↓  ↓  ↓
[10][20][30][0][0]
```

### 4. Stack Memory Allocation

Arrays are typically allocated on the stack (unless they're part of a heap-allocated structure):
- Faster allocation and deallocation
- Automatic memory management
- Limited by stack size

## Complete Example Output

```go
func main() {
    // Method 1
    var arr1 [5]int
    fmt.Println(arr1) // [0 0 0 0 0]
    
    var arr2 [3]string
    fmt.Printf("%q\n", arr2) // ["" "" ""]
    
    arr1[0] = 10
    arr1[1] = 20
    arr1[2] = 30
    
    arr2[0] = "Hello"
    arr2[1] = "World"
    
    fmt.Println(arr1)        // [10 20 30 0 0]
    fmt.Printf("%q\n", arr2) // ["Hello" "World" ""]
    
    // Method 2
    arr3 := [6]int{1, 2, 3, 4, 5, 6}
    fmt.Println(arr3) // [1 2 3 4 5 6]
    
    // Method 3
    arr4 := [...]string{"Go", "is", "awesome"}
    fmt.Printf("%q\n", arr4) // ["Go" "is" "awesome"]
    
    // Method 4
    var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
    fmt.Println(arr) // [[1 2 3] [4 5 6]]
    
    // Value type demonstration
    arr5 := arr3
    arr5[0] = 100
    fmt.Println("arr3:", arr3) // arr3: [1 2 3 4 5 6]
    fmt.Println("arr5:", arr5) // arr5: [100 2 3 4 5 6]
    
    // Length
    fmt.Println("Length of arr4:", len(arr4)) // Length of arr4: 3
}
```

## Arrays vs Slices

| Feature | Arrays | Slices |
|---------|--------|--------|
| Size | Fixed at compile time | Dynamic |
| Type | Size is part of type | Size not part of type |
| Memory | Value type (copied) | Reference type |
| Performance | Faster for small, fixed data | Better for dynamic data |
| Common usage | Less common | More common |

## Best Practices

### ✅ DO:

1. **Use arrays when the size is known and fixed**
   ```go
   var days [7]string = [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
   ```

2. **Use ellipsis for readability when size is obvious**
   ```go
   primes := [...]int{2, 3, 5, 7, 11, 13}
   ```

3. **Use arrays for small, fixed-size data structures**
   ```go
   type RGB [3]uint8
   color := RGB{255, 128, 0}
   ```

### ❌ DON'T:

1. **Don't use large arrays as function parameters (pass pointers or use slices instead)**
   ```go
   // BAD - copies entire array
   func processArray(arr [1000000]int) { }
   
   // GOOD - uses pointer
   func processArray(arr *[1000000]int) { }
   
   // BETTER - use slice
   func processArray(arr []int) { }
   ```

2. **Don't use arrays when you need dynamic sizing**
   ```go
   // BAD - size must be constant
   // var n = getUserInput()
   // var arr [n]int // Won't compile
   
   // GOOD - use slice
   var n = getUserInput()
   arr := make([]int, n)
   ```

## Common Operations

### Iterating Over Arrays

**Using traditional for loop:**
```go
for i := 0; i < len(arr3); i++ {
    fmt.Println(arr3[i])
}
```

**Using range:**
```go
for index, value := range arr3 {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}
```

**Only values:**
```go
for _, value := range arr3 {
    fmt.Println(value)
}
```

### Comparing Arrays

Arrays of the same type can be compared using `==` and `!=`:

```go
arr1 := [3]int{1, 2, 3}
arr2 := [3]int{1, 2, 3}
arr3 := [3]int{1, 2, 4}

fmt.Println(arr1 == arr2) // true
fmt.Println(arr1 == arr3) // false
```

## Key Takeaways

- Arrays have **fixed size** defined at compile time
- Array size is **part of the type**: `[5]int` ≠ `[10]int`
- Arrays are **value types** - assignments create copies
- Stored in **contiguous memory** for efficient access
- Use **slices** for most dynamic array operations
- Use **arrays** when size is fixed and known at compile time

## When to Use Arrays

✅ Fixed-size collections (days of week, RGB colors)  
✅ Performance-critical code with known sizes  
✅ When you need value semantics (copies)  
✅ Embedded systems with memory constraints  

❌ Dynamic collections  
❌ Function parameters (use slices instead)  
❌ When size varies at runtime  

## Further Reading

- **Slices**: The more commonly used dynamic alternative to arrays
- **Array pointers**: When you need to pass arrays efficiently
- **Memory layout**: Understanding how arrays affect performance

## Usage

Simply run the program:

```bash
go run main.go
```

