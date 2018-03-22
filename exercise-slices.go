/*
Exercise: Slices

Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)
*/

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	finalVal := make([][]uint8, dy, dx*dy)
	for Y := 0; Y < dy; Y++ {
		// using equation (x + y)/2, where x = 1 && y = 1
		tempVal := make([]uint8, dx)
		for X := 0; X < dx; X++ {
			pixel := ((1 + (1 * dx)) + (1 + (1 * dy))) / 2
			tempVal[X] = uint8(pixel)
		}

		finalVal[Y] = tempVal
	}

	return finalVal
}

func main() {
	pic.Show(Pic)
}
