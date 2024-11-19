package main

import (
	"errors"
	"fmt"
)

func main() {

	// var sum, err = sum(69, 420)
	var sum, err = sum(666, 420)

	if err != nil {
		fmt.Printf("Not nice!\n%s", err)
		return
	}

	fmt.Printf("The nice number is %d", sum)
}

func sum(a, b int) (int, error) {
	thesum := a + b

	if thesum >= 69 && thesum <= 666 {
		return thesum, nil
	}

	return -1, errors.New("Too big!")
}
