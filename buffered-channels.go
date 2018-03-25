/*
Buffered Channels

Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

ch := make(chan int, 100)

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Modify the example to overfill the buffer and see what happens.
*/

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	for i := 0; i < 1000; i++ {
		ch <- i
	}

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
