package main

import (
	"fmt"
	"io"
)

type MyWreader []byte

func (M *MyWreader) Write(p []byte) (int, error) {
	*M = p
	return len(p), nil
}

func (M *MyWreader) Read(buf []byte) (int, error) {
	buf = *M
	fmt.Println(string(buf))
	return len(buf), io.EOF
}
func main() {

	var b = new(MyWreader)
	io.WriteString(b, "hello")
	io.ReadAll(b)

}
