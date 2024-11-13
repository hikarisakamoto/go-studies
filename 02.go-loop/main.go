package main

import "fmt"

func main() {
	var array [3]int
	array[0] = 420
	array[1] = 69
	array[2] = 666

	for index, value := range array {
		fmt.Printf("This is pos [%d] with value [%d].\n", index, value)
	}

	fmt.Printf("This array has a size of %d", len(array))
}
