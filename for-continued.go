/*
For Continued

The init and post statement are optional.
*/

package main

import "fmt"

func main() {
	sum := 1
	for sum < 100 { // equivalent to a while() statement in C
		sum += sum
	}

	fmt.Println(sum)
}
