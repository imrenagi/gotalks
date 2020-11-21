package main

import "fmt"

func main() {
	c := make(chan bool)
	c <- true

	fmt.Println("this line will never be printed")
}
