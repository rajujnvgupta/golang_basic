package main
import "fmt"
import "strconv"
import "reflect" // The reflect library is used to obtain the datatype of an object.

func main() {
  num1:= 10
  fmt.Println(num1, reflect.TypeOf(num1))
  str:= strconv.Itoa(num1)
  fmt.Println(str, reflect.TypeOf(str))
}
