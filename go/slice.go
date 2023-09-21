package main

import "fmt"

func main() {
	t := "asdf"
    user := []string{t}
	user = append(user, "raju1")
	user = append(user, "raju2")

	for _, use := range user {
		fmt.Println(use)
	}
}
