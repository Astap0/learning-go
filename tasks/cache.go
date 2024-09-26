package main

import "fmt"

type User struct{
	name string
	age int
}

func NewUser(name string, age int) User{
	return User{
		name: name,
		age: age,
	}
}
func (u User) isAdult() bool{
	if u.age>=18{
		return true
	} else {
		return false
	}
} 

func (u User) printUserInfo() {
	fmt.Printf("Имя: %s, Возвраст: %d", u.name, u.age)
}
func main(){
	user:=NewUser("Sasha",42)
	fmt.Println(user.isAdult())
	user.printUserInfo()
}