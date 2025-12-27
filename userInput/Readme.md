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
