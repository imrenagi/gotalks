package main

import (
	"fmt"
	"time"
)

func main() {
	//START, OMIT
	for _, v := range []int{1, 2, 3, 4, 5} {
		go func(i int) { // HLxx
			//can do anything else
			fmt.Println(i)
		}(v) // HLxx
	}
	//STOP, OMIT

	<-time.After(2 * time.Second)
}
