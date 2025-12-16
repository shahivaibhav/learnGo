package main

import (
	"fmt"
	myutils "go_tutorials/myUtils"
	variables "go_tutorials/variablesDemo"
)

func main() {
	fmt.Println("Hello From Vaibhav!")
	myutils.PrintHello()

	fmt.Println("----- Variables Demo -----")
	variables.PrintVariables()
	variables.ExportedFunction()
}
