package main

import "fmt"

func main(){

	//this code failed
	channel := make(chan string)

	channel <- "first message"
//	go func(){
//		channel <- "first message"
//	}()

	fmt.Println(<-channel)
}
