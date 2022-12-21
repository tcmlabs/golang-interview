package main

import (
	"fmt"
	"time"
)

type User struct {
	ID int
}

func (s *User) ChangeID() {
	s.ID += 1
}

func (s *User) GetID() {
	fmt.Println(s.ID)
}

func main() {
	// Preparing the data set
	var ids []int
	for i := 1; i < 100; i++ {
		ids = append(ids, i)
	}
	u := User{
		ID: 0,
	}
	for range ids {
		go u.ChangeID()
	}
	time.Sleep(time.Second)

	// What is the result of this command? Why?
	u.GetID()

	// Fix the behavior of this function by making every go routine increment the user.ID one by one
}
