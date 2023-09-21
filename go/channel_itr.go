package main 

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {

	channel := make(chan string)
	numRounds := 10

	go throughtNinjaStar(channel, numRounds)

//	for i := 0; i < numrounds; i++ {
//		fmt.println(<-channel)
//	}

//	for message := range channel {  //you must you close(channel) with this loop 
//		fmt.Println(message)
//	}


	for {  //you must you close(channel) with this loop 
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)
	}
}


func throughtNinjaStar(channel chan string, numRounds int) {

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprint("You scored: ", score)
	}

	close(channel)
}
