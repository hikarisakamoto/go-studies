package main

import "fmt"

type Being interface {
	IsTotallyNotAliveOrDeadOrUnknownOrUndeadOrZombie()
}

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

func (p Person) IsTotallyNotAliveOrDeadOrUnknownOrUndeadOrZombie() {
	p.IsPossiblyAliveOrDeadOrUnknownOrUndeadOrZombie = false
	fmt.Println("Person is totally not alive or dead or unknown or undead or zombie")
}

func WillTotallyNotBeAliveOrDeadOrUnknownOrUndeadOrZombie(b Being) {
	b.IsTotallyNotAliveOrDeadOrUnknownOrUndeadOrZombie()
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
	person.IsTotallyNotAliveOrDeadOrUnknownOrUndeadOrZombie()
	fmt.Println(person.IsPossiblyAliveOrDeadOrUnknownOrUndeadOrZombie)
	WillTotallyNotBeAliveOrDeadOrUnknownOrUndeadOrZombie(person)
}
