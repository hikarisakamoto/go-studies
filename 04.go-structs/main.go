package main

import "fmt"

type Stats struct {
	HP                                           int
	MP                                           int
	HasSuperPowersOrIsAMutantOrIsAnAlienOrIsAGod bool
}

type Person struct {
	Name                                           string
	Age                                            int
	IsPossiblyAliveOrDeadOrUnknownOrUndeadOrZombie bool
	Stats
}

func main() {
	// Create a new instance of Stats
	stats := Stats{HP: 100, MP: 50, HasSuperPowersOrIsAMutantOrIsAnAlienOrIsAGod: true}
	fmt.Println(stats)

	// Create a new instance of Person
	person := Person{Name: "John Doe", Age: 30, IsPossiblyAliveOrDeadOrUnknownOrUndeadOrZombie: true, Stats: stats}
	fmt.Println(person)

	// Accessing fields
	fmt.Println(person.Stats.HP)
	fmt.Println(person.HP)
}
