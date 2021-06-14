package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (m MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func Readers() {
	reader.Validate(MyReader{})
}
