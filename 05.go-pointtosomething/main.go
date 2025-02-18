package main

func main() {

	a := 42
	println(a)

	var pointer *int = &a
	println(pointer)
	*pointer = 69
	println(a)

	b := &a
	println(b)

	*b = 96
	println(a)
}
