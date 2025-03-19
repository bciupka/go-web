package main

import "fmt"

type User struct {
	FirstName string
	LastName string
	Age int
}

func (u *User) DoubleAge() {
	u.Age *= 2
}

func main() {
	user := User {
		FirstName: "Bartek",
		LastName: "C",
	}

	fmt.Println(user.FirstName, user.LastName, user.Age)

	user2 := User{
		FirstName: "Aga",
		LastName: "D",
		Age: 28,
	}

	fmt.Println(user2.FirstName, user2.LastName, user2.Age)

	user2.DoubleAge()
	fmt.Println(user2.Age)
}