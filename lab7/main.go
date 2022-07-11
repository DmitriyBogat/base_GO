package main

import (
	"fmt"
	"io"
	"strings"
)

type MyWreader []byte

func (M *MyWreader) Write(p []byte) (int, error) {
	*M = p
	return len(p), nil
}

func (M *MyWreader) Read(p []byte) (int, error) {
	s := make([]byte, len(*M))
	buf := strings.NewReader(string(*M))
	n, err := buf.Read(s)
	fmt.Println(string(s))
	fmt.Printf("Прочитано %v байт", n)
	if err != nil {
		fmt.Println(err)
	}
	return len(p), nil
}
func main() {

	var b = new(MyWreader)
	io.WriteString(b, "hello")
	io.ReadAll(b)

}
