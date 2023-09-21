package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	now := time.Now()
	p(now)
	then := time.Date(2019, 11, 17, 20, 34, 58, 2234455, time.UTC)
	p(then)
}
