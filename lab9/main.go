package main

import (
	"fmt"
	"main/lab9/flags"
)

func main() {
	var F flags.Flags
	var fileF flags.Flags
	F = flags.GetFlag()
	fileF = flags.GetFileflag()
	fmt.Println(F.Log)
	fmt.Println(F.Sentry_url)
	fmt.Println(fileF.Log)
}
