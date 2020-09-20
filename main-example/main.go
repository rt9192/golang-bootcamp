package main

import "fmt"

func main() {
	fmt.Println("Hello Golang !!")
	Hye()

	// Read the below comments for guidelines to functions importable from a package
	bye()
	// Importable only in the same package
	// Functions with names that have first letter as CAPS are importable into other packages
}
