package userinput

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() {
	fmt.Println("-----Taking user input using Scan-----")

	// Scan function is used to take input from the user.
	// It reads input until a newline is encountered.
	// It can read multiple values separated by spaces.
	// The input values need to be passed as pointers to the Scan function.
	// Scan is useful for simple input scenarios where you want to read values directly into variables.
	// Scan is often used in console applications for interactive user input.
	// Scan returns the number of items successfully scanned and an error if any occurred during scanning.

	fmt.Println("Enter your name:")
	var (
		first_name string
		last_name  string
	)
	fmt.Scan(&first_name, &last_name)
	fmt.Println("Hello, " + first_name + " " + last_name + "!")

	//But if you want to read a full line of input including spaces, you can use Scanln or Scanf.
	// The key difference is that Scanln reads input until a newline, while Scanf allows for formatted input.

	fmt.Println("-----Taking user input using Scanln-----")
	fmt.Println("Enter your full name:")
	var full_name string
	fmt.Scanln(&full_name)
	fmt.Println("Hello, " + full_name + "!")

	fmt.Println("-----Taking user input using Scanf-----")

	fmt.Println("Enter your age and city (e.g., 25 NewYork):")
	var (
		age  int
		city string
	)
	fmt.Scanf("%d %s", &age, &city)
	fmt.Printf("You are %d years old and live in %s.\n", age, city)
}

func UserInputUsingBuffIo() {
	//There is a standard way for taking input from the user using bufio
	// This method is more efficient and flexible for handling larger inputs or more complex input scenarios.
	// It allows for buffered reading, which can improve performance when dealing with large amounts of data.
	// It also provides more control over how input is read and processed by using various methods available in the bufio package.
	// This method is particularly useful in scenarios where you need to read input line by line or when dealing with files.
	// this is because it minimizes the number of read operations by buffering the input.
	// Buffering means that instead of reading one byte or character at a time, larger chunks of data are read into memory at once.
	// This reduces the overhead of multiple read calls, which can be costly in terms of performance, especially when reading from slower input sources like files or network connections.

	fmt.Println("-----Taking user input using Buffio-----")

	input := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your address:")
	adress, _ := input.ReadString('\n')
	fmt.Println("Your address is: " + adress)

	//Must know Methods of Bufio:
	//1. NewReader: This function creates a new buffered reader that reads from the specified io.Reader (in this case, os.Stdin for standard input).
	//2.
}
