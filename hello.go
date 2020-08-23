package main

import "fmt"

func main() {
	fmt.Println("Go, world!")

	var buffer [256]int
	fmt.Printf("The buffer has %d elements.\n", len(buffer))

}
