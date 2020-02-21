package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (mr MyReader) Read(b []byte) (n int, er error) {
	for i := 0; i < len(b); i++ {
		b[i] = 65
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
