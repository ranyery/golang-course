package main

import "fmt"

func main() {
	var arr [3]int // Indica que o array terá 3 posições fixas
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	// arr[3] = 4 => // Resulta em erro

	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Printf("O último valor do array é %v\n", arr[len(arr)-1])

	for i, value := range arr {
		fmt.Printf("O índice é %d e o valor é %d\n", i, value)
	}
}
