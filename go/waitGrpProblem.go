package main 


import (
	"fmt"
	"sync"
)

func main(){

	//indefinity wait 
	var beeper sync.WaitGroup
	fmt.Println("xyz")
	beeper.Add(1)
	beeper.Wait()

}

