package main

import "fmt"

func main() {
	// Various ways to declare an array in Go

	// Method 1: Declare an array with a fixed size

	var arr1 [5]int   //Tell size of array and type of array
	fmt.Println(arr1) // By deafult as the type of array is int all values are 0

	var arr2 [3]string       // String array of size 3
	fmt.Printf("%q\n", arr2) // By default all values are empty strings ("")

	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30

	arr2[0] = "Hello"
	arr2[1] = "World"

	fmt.Println(arr1)
	fmt.Printf("%q\n", arr2)

	// Method 2: Declare and initialize an array in one line

	arr3 := [6]int{1, 2, 3, 4, 5, 6} //Use curly braces to initialize
	fmt.Println(arr3)

	// Method 3: Let Go determine the size of the array
	// by using "..." instead of a size
	// This is called as ellipsis

	arr4 := [...]string{"Go", "is", "awesome"}
	fmt.Printf("%q\n", arr4)

	// Method 4: Multidimensional arrays(2D arrays)

	var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}} //2 rows and 3 columns
	fmt.Println(arr)

	// Important concepts to know about arrays in Go.

	// 1. Arrays are value types in Go.
	// This means when you assign an array to another array,
	// a copy of the original array is made.

	arr5 := arr3 // arr5 is a copy of arr3
	arr5[0] = 100
	fmt.Println("arr3:", arr3) // arr3 remains unchanged
	fmt.Println("arr5:", arr5) // arr5 has the modified value

	// 2. Length of an array can be obtained using len() function
	fmt.Println("Length of arr4:", len(arr4))

	//3. They are stored in contiguous memory locations
	// This allows for efficient access and manipulation of array elements.

	//4. Arrays are stored in stack memory

}
