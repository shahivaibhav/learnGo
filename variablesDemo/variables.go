package variables

import "fmt"

//What is difference between var and const?
//Var is used to declare variables whose values can be changed later in the program.
//Const is used to declare constants whose values cannot be changed once assigned.

//What are different ways to declare variables in Go?
//1. Using var keyword with explicit type
//2. Using var keyword with type inference
//3. Short variable declaration using :=

// Examples:

// Explicit type declaration
var name string = "Vaibhav"
var age int = 25
var isDeveloper bool = true

const country string = "India"
const state string = "Haryana"
const salary float64 = 75000.50

// Can we declare multiple variables at once in Go? If yes, how?
// Yes, we can declare multiple variables at once using var block or multiple assignment.

var (
	city        string = "Gurgaon"
	countryCode int    = 91
)

// Type inference example
const (
	pi       = 3.14
	language = "Go"
)

func PrintVariables() {
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Developer:", isDeveloper)
	fmt.Println("Country:", country)
	fmt.Println("State:", state)
	fmt.Println("Salary:", salary)
	fmt.Println("City:", city)
	fmt.Println("Country Code:", countryCode)
	fmt.Println("Pi:", pi)
	fmt.Println("Language:", language)

	// Why take variable declaration using := is not allowed at package level?
	// The short variable declaration using := is only allowed within functions.
	// At the package level, variables must be declared using the var keyword or const keyword.
	take := 19
	fmt.Println("Take:", take)

}

// To export a variable or function in Go, the variable name must start with an uppercase letter.
func ExportedFunction() {
	fmt.Println("This is an exported function.")
}
