package functions

//First Syntax of function in Go
//func functionName(parameters) returnType {
//	function body
//	return statement
//}

// Example 1: Function with no parameters and no return type
func addTwoNumbers() {
	sum := 5 + 10
	println("Sum is:", sum)
}

// Example 2: Function with parameters and no return type
func addNumbers(a int, b int) {
	res := a + b
	println("Result is:", res)
}

// Example 3: Function with parameters and return type
func multiplyNumbers(a int, b int) int {
	return a * b
}

//Example 4 : Better way to declare function with multiple parameters of same type

func betterMultiplyNumbers(a, b int) int {
	return a * b
}

// Example 5: Function with multiple return values
func swapValues(a string, b string) (string, string) {
	return b, a
}

// Example 6: Function with named return values
func divideNumbers(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

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
