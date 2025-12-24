package underscoreidentifier

import "fmt"

// In Go, an underscore (_) is known as the blank identifier.
// It is used to ignore values that are not needed.
// The blank identifier can be used in various scenarios, such as when you want to ignore a return value from a function or when you want to ignore an index in a loop.

// Example 1: Ignoring return values from a function

func divideTwoNumbers(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}

	return a / b, nil
}

// Example 2: Ignoring index in a loop

func ignoreIndexInLoop() {
	fmt.Println("----- Ignoring Index in Loop Demo -----")
	numbers := []int{10, 20, 30, 40, 50}
	sum := 0

	for _, value := range numbers {
		sum += value
	}

	fmt.Println("Sum of numbers is:", sum)
}

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
