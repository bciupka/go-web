package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}
type MyReaderError int

func (e MyReaderError) Error() string {
	return fmt.Sprintf("%v capacity is inacceptable", int(e))
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (r MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return -1, MyReaderError(cap(b))
	}
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
