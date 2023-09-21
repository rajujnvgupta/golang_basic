package main
import (
	"fmt"
	"time"
	"os"
)


func fibo(n int) int {

	if n == 0 || n == 1{
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
	time.Sleep(5*time.Second)

	print(fibo(7))
}

func createFile(p string) *os.File {

	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File){

	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File){

	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}


