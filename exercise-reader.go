/*
Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
*/

package main

import "golang.org/x/tour/reader"

type MyReader struct {}

func (r MyReader) Read(b []byte) (n int, err error) {
	b = b[:cap(b)]
	for i := range b {
		b[i] = 'A'
	}
	return cap(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
