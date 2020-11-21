package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(15123)
}

// START, OMIT
func main() {
	c := make(chan int)

	go func() {
		<-time.After(time.Duration(rand.Intn(2)) * time.Second)
		c <- 10
	}()

	select { // HL12
	case val := <-c: // HL12
		fmt.Println(val)
	case <-time.After(time.Duration(rand.Intn(2)) * time.Second): // HL12
		fmt.Println("timeout")
	}
}

// STOP, OMIT
