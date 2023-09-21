package main

import (
	"fmt"
	"time"
)

func main(){
	timer1 := time.NewTimer(2*time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func(){
		fmt.Println("before timer2 block")
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()

	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("timer 2 stopped")
		fmt.Println(time.Second)
	}

	time.Sleep(2*time.Second)
}
