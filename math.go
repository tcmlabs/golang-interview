package main

import (
	"fmt"
	"os"
	"reflect"
)

/////////////////////////////////////////////////////////
type monoidElement interface {
	int | bool | string | []string
}
type operation[T monoidElement] func(T, T) T

type monoid[T monoidElement] struct {
	name      string
	element   T
	operation operation[T]
	neutral   T
}

/////////////////////////////////////////////////////////

// Example:
var integerMonoidOne = monoid[int]{
	name:      "integer monoid 1",
	element:   42,
	operation: func(a, b int) int { return a + b },
	neutral:   0,
}

// Use another operation than addition
var integerMonoidTwo = monoid[int]{
	name:      "integer monoid 2",
	element:   42,
	operation: func(a, b int) int {},
	neutral:   nil,
}

// String monoid
var stringMonoid = monoid[string]{
	name:      "string monoid",
	element:   "Hello World",
	operation: func(a, b string) string {},
	neutral:   nil,
}

// List monoid
var listMonoid = monoid[[]string]{
	name:      "list monoid",
	element:   []string{"Hello, World"},
	operation: func(a, b []string) []string {},
	neutral:   nil,
}

// Bool monoid AND
var boolMonoidAnd = monoid[bool]{
	name:      "bool monoid AND",
	element:   true,
	operation: func(a, b bool) bool { return a && b },
	neutral:   nil,
}

// Bool monoid OR
var boolMonoidOr = monoid[bool]{
	name:      "bool monoid OR",
	element:   false,
	operation: func(a, b bool) bool { return a || b },
	neutral:   nil,
}

/////////////////////////////////////////////////////////
func monoidRun[T monoidElement](monoidObj monoid[T]) bool {
	return reflect.DeepEqual(monoidObj.operation(monoidObj.element, monoidObj.neutral), monoidObj.element)
}

func assertMonoid[T monoidElement](monoidObj monoid[T]) {
	fmt.Printf("=> Monoid: %s ", monoidObj.name)
	if !monoidRun(monoidObj) {
		fmt.Println(": FAILED")
		os.Exit(1)
	}
	fmt.Println(": PASSED")
}

func main() {
	assertMonoid(integerMonoidOne)
	assertMonoid(integerMonoidTwo)
	assertMonoid(stringMonoid)
	assertMonoid(listMonoid)
	assertMonoid(boolMonoidAnd)
	assertMonoid(boolMonoidOr)
}
