package main

import (
	"fmt"
	"go_tutorials/functions"
	myutils "go_tutorials/myUtils"
	underscoreidentifier "go_tutorials/underScoreIdentifier"
	userinput "go_tutorials/userInput"
	variables "go_tutorials/variablesDemo"
)

func PrintlnUseCase() {
	// Println function directly adds a newline at the end.
	// Println need not take format specifiers.
	// It automatically adds spaces between arguments.
	// Println is useful for quick and simple output.
	// Println is often used for debugging or logging simple messages.
	// Println is straightforward and easy to use for basic output needs.
	// Println is commonly used in scenarios where readability and simplicity are prioritized over precise formatting.

	fmt.Println("----------------------Println Use Case----------------------")
	const (
		name       = "Vaibhav"
		age        = 25
		salary     = 50000.50
		start_char = 'V'
	)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Start Character:", start_char)

	fmt.Println("----------------------")

	fmt.Println(name, age, salary, start_char)
}

func PrintfUseCase() {
	// Printf function does not add a newline at the end.
	// Printf requires format specifiers to format the output.
	// It does not add spaces between arguments automatically.
	// Printf is useful when you need precise control over the output format.
	// Printf is often used for formatted strings, such as when printing numbers with specific decimal places.
	// Printf is also used when you want to create a formatted string without printing it immediately, using fmt.Sprintf.
	// Printf is more versatile for complex formatting needs.
	// Printf is commonly used in scenarios where the output format is critical, such as in logging or generating reports.
	fmt.Println("----------------------Printf Use Case----------------------")
	const (
		name       = "Vaibhav"
		age        = 25
		salary     = 50000.50
		start_char = 'V'
	)

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Salary : %.3f\n", salary)
	fmt.Printf("Start Character: %c\n", start_char)

	fmt.Printf("----------------------\n")

	fmt.Printf("Combined Output without spaces:\n")
	fmt.Printf("%s%d%.2f%c\n", name, age, salary, start_char)
	fmt.Printf("Combined Output with spaces:\n")
	fmt.Printf("%s %d %.2f %c", name, age, salary, start_char)
}

func main() {
	fmt.Println("Hello From Vaibhav!")
	myutils.PrintHello()

	fmt.Println("----- Variables Demo -----")
	variables.PrintVariables()
	variables.ExportedFunction()

	fmt.Println("Difference between Println and Printf")
	PrintlnUseCase()
	PrintfUseCase()

	fmt.Println("----- User Input Demo -----")
	userinput.UserInput()
	userinput.UserInputUsingBuffIo()
	functions.LearnFunctions()
	underscoreidentifier.UnderscoreIdentifierDemo()
}
