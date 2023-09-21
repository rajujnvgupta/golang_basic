package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key int 
	resp chan int
}
type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {

	var readOp uint64
	var writeOp uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func(){
			for {
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int)
				}

				reads <- read
				<-read.resp
				atomic.AddUnit64(&readOp, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; i < 10; w++ {
		go func() {
			for {
				write := writeOp {
					key: rand.Intn(5)
					val: rand.Intn(100)
					resp: make(chan bool)
				}

				writes <- write 
				<-write.resp
				atomic.AddUnit64(&writeOp, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUnit64(&readOp)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUnit64(&writeOp)
	fmt.Println("writeOp:", writeOpsFinal)
}
