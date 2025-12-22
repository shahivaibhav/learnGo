package userinput

import "fmt"

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
