/*
Functions

A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.
*/

package main

import "fmt"

// Notice that the variable type comes after the variable name, and that the
// return type comes after the function definition.
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
