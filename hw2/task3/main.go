package main

import (
	reader "golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (reader MyReader) Read(b []byte) (int, error) {
	n := len(b)
	for i := 0; i < n; i++ {
		b[i] = 'A'
	}
	return n, nil
}

func main() {
	reader.Validate(MyReader{})
}
