package main

import "fmt"

func main() {
	//START, OMIT
	c := make(chan bool, 5) // HL
	c <- true

	fmt.Println("this line be printed")
	//START, OMIT
}
