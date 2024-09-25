package main

import "fmt"

func main(){
	defer handerPanic()

	var messages []string
	for i:=1; i<=4; i++ {
		message:=fmt.Sprintf("Message #%d", i)
		messages = append(messages, message)		
	}
	fmt.Println(messages)
	messages[4] = "Message #5"
	fmt.Println(messages)
}

func handerPanic(){
	if r := recover(); r!= nil {
		fmt.Println("Recovered from panic:",r)
	}
}