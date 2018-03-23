/*
Exercise: Fibonacci closure

Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/

package main

import "fmt"

func fibonacci() func() int {
	fibCount := 0
	fibPrev  := 0
	
	fibFunc := func() int {
		// [ NOTE ] cfarvin: May be a good place to use defer, though both work.
		fibReturn := fibCount + fibPrev
		fibPrev = fibCount
		fibCount++
		return fibReturn
	}
	
	return fibFunc
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
