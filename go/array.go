package main

import "fmt"

func main() {

    var twoD [10][10]int
    for i := 0; i < 10; i++ {
        for j := 0; j < 10; j++ {
            twoD[i][j] = i + j
        }
    }
	for i:=0; i < len(twoD); i++{
		fmt.Println(twoD[i])
	}
}
//[0 1 2]
//[1 2 3]
//[2 3 4]
