package main

import "fmt"

func mayPanic() {
	panic("a problem raju")
}

func main(){

	defer func(){
		if r := recover(); r != nil{
			fmt.Println("Recoverd. Error:\n", r)
		}
	}()

	fmt.Println("Raju before mayPanic")
	mayPanic()

	fmt.Println("After mayPanic")
}
