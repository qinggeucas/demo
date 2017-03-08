package main

import "fmt"

type User struct {
	Username string
	Password string
}

func (this *User) Say(word string) {
	fmt.Println(fmt.Sprintf("%s say %s", this.Username, word))
}
