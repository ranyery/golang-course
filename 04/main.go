package main

import "fmt"

type ID int

var id ID

func main() {
	fmt.Printf("O tipo de id é %T\n", id)
	fmt.Printf("O valor de id é %v", id)
}
