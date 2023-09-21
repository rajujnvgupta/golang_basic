package main

import (
	"fmt"
	"os"
)

func main(){

	panic("a problem")

	_, err := os.Create("/tmp/file")
	fmt.Println("tmp raju")
	if err != nil {
		fmt.Println("tmp raju")
		panic(err)
	}
}
