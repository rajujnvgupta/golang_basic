package main

import "fmt"

type sumInterface interface {
	sum() int
}

//structure

type values struct {
	first int
	second int
}

func (v values) sum() int {
	return v.first + v.second
}


//func main() {
//	var v sumInterface 
//	v = &values{first: 1, second: 2}
//	fmt.Println(v.sum())
//}



//func main() {
//	var v values
//	v.first = 1
//	v.second = 3
//	fmt.Println(v.sum())
//}


//func main() {
//	testmap := make(map[int]int)
//	testmap[0] = 10
//
//	fmt.Println(testmap)
//
//	for k, v := range testmap{
//		fmt.Println(k, v)
//	}
//}


//func main() {
//	var test interface{}
//
//	test = "some string"
//	print(test)
//
//	test = 10
//	print(test)
//}
//
//func print(t interface{}){
//	fmt.Println(t)
//}



func main() {
	test := []interface{}{"raju", 23}
	test = append(test, 1234)

	print(test)
}

func print(t interface{}){
	fmt.Println(t)
}

