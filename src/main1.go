package main

import (
	"fmt"
	"github.com/kmlasharma/CS4098/src/pml"
	"os"
)

func main() {
	reader, _ := os.Open("./pml/test.pml")
	parser := pml.NewParser(reader)
	process := parser.Parse()
	fmt.Println("Process: ", process)
	str := process.Encode()
	fmt.Println("Str:\n", str)
}
