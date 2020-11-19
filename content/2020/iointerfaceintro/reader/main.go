package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	//STARTMAIN1, OMIT

	reader := strings.NewReader("Hello my name is Imre Nagi")
	p := make([]byte, 4)

	for {
		n, err := reader.Read(p)
		if err != nil {
			if err == io.EOF {
				fmt.Println(string(p[:n]))
				break
			}
			fmt.Println(err)
			return
		}
		fmt.Println(string(p[:n]))
	}

	//STOPMAIN1, OMIT
}
