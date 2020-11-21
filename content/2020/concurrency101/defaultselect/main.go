package main

import (
	"fmt"
	"time"
)

//START, OMIT
func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default: // HL13
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

//STOP, OMIT
