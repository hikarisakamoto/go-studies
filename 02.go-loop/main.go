package main

import "fmt"

func main() {
	var array [4]int
	array[0] = 420
	array[1] = 69
	array[2] = 666

	for index, value := range array {
		fmt.Printf("This is pos [%d] with value [%d].\n", index, value)
	}

	fmt.Printf("This array has a size of %d", len(array))

	fmt.Printf("\n----------------------------------------\n")

	var slice = []int{69, 420, 666, 42069, 69420}

	fmt.Printf("len=%d, cap=%d values=%v\n", len(slice), cap(slice), slice)

	fmt.Printf("[:0]len=%d, [:0]cap=%d [:0]values=%v\n", len(slice[:0]), cap(slice[:0]), slice[:0])

	fmt.Printf("[:4]len=%d, [:4]cap=%d [:4]values=%v\n", len(slice[:4]), cap(slice[:4]), slice[:4])

	fmt.Printf("[2:]len=%d, [2:]cap=%d [2:]values=%v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	nice := 96
	slice = append(slice, nice)

	fmt.Printf("Appended %d to slice\n", nice)
	fmt.Printf("len=%d, cap=%d values=%v\n", len(slice), cap(slice), slice)
}
