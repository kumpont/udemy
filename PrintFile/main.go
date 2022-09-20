package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	filename := os.Args[1]
	myfile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, myfile)
}
