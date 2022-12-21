package main

import "fmt"

// Entity
/*
Definition of an Entity:
 - Speed() get the current entity speed
 - Walk() increase the speed slowly
 - Run() increase the speed quickly
*/
type Entity interface {
	Speed() int
	Walk()
	Run()
}

type Human struct {
	speed int
}

type Dog struct {
	speed int
}

// Increase the speed of an entity
func action(entity Entity) {
	fmt.Printf("Current speed: %d...\n", entity.Speed())
	fmt.Printf("Starting to walk:...\n")
	entity.Walk()
	fmt.Printf("Now walking at %d k/h:...\n", entity.Speed())
	fmt.Printf("Starting to run:...\n")
	entity.Run()
	fmt.Printf("Now running at %d k/h:...\n", entity.Speed())
}

func main() {
	human := Human{speed: 0}
	dog := Dog{speed: 0}

	// Make the objects able to perform the action
	action(human)
	action(dog)
}
