package main

import "fmt"

func main() {
	c := make(chan bool)
	go func() {
		<-c
	}()
	c <- true

	fmt.Println("this line is printed")
}
