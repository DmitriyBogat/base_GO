package main

import (
	"fmt"
	"main/lab8/flags"
)

func main() {
	var F flags.Flags
	F = flags.GetFlag()
	fmt.Println(F.Log)
	fmt.Println(F.Jaeger_url)
}
