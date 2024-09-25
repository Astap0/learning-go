package main

import "fmt"

type Age int

type User struct{
	name	string
	age 	Age
	sex		string
	weight	int
	height	int
}

func (a Age) isAdult() bool{
	return a>18
}


func NewUser(name,sex string, age Age, weight, height int) User{
	return User{
		name:	name,
		age:	age,
		sex:	sex,
		weight:	weight,
		height:	height,
	}
}

func (u User) printUserInfo() {
	fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
}

func main() {
user1:=User{"Vasya", 21, "Male", 75, 185}
user2:=User{"Sergey", 23, "Male", 83, 200}
fmt.Printf("%+v\n", user1)
fmt.Printf("%+v\n", user2)
user1.printUserInfo()
fmt.Println(user1.age.isAdult())
}

// // В Go есть два основных типа методов-получателей (receivers), которые можно использовать в структуре:

// // Value receiver (Получатель по значению).
// func (u User) printUserInfo() {
//     fmt.Println(u.name, u.age, u.sex, u.weight, u.height)
// }
// // Pointer receiver (Получатель по указателю).
// func (u *User) updateWeight(newWeight int) {
//     u.weight = newWeight
// }