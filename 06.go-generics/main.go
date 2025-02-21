package main

import "fmt"

type Tilda int

// Number is a constraint that represents int or float64 type.
type Number interface {
	~int | ~float64
}

func Sum[T Number](value map[string]T) T {
	var sum T
	for _, v := range value {
		sum += v
	}
	return sum
}

func main() {
	intMap := map[string]int{"A": 10, "B": 20, "C": 30}
	floatMap := map[string]float64{"A": 10.5, "B": 20.5, "C": 30.5}
	tilda := map[string]Tilda{"A": 10, "B": 20, "C": 30}

	fmt.Println(Sum(intMap))   // 60
	fmt.Println(Sum(floatMap)) // 61.5
	fmt.Println(Sum(tilda))    // 60
}
