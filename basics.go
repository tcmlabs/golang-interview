package main

import (
	"encoding/json"
	"fmt"
)

type MyObj struct {
	Name string
	Age  int
}

func changeMe(str1 *string, str2 string) {
	*str1 = "You've changed"
	str2 = "You've changed"
}

func amIDefering(size int) {
	for i := 1; i <= size; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	// 1 - What will be displayed? Why?
	one := "one"
	two := "two"
	changeMe(&one, two)
	fmt.Printf("One: %s, Two: %s\n", one, two)
	return // Remove me once question answered

	// 2 - What will be displayed? Why?
	amIDefering(5)
	return // Remove me once question answered

	// 3 - What will de displayed? Fix it without removing any code, only adding code.
	myJson := []byte("{\"myname\":\"Test\",\"myage\":42}")
	var myObj MyObj
	_ = json.Unmarshal(myJson, &myObj)
	fmt.Printf("%+v", myObj)
	return // Remove me once question answered

}
