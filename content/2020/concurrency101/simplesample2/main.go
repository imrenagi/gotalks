package main

import (
	"fmt"
	"time"
)

//START, OMIT
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
}

//STOP, OMIT
