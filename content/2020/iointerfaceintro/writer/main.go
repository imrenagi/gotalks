package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	//STARTMAIN1, OMIT
	names := []string{
		"Imre Nagi",
		"Foo Bar",
	}
	var writer bytes.Buffer

	for _, name := range names {
		n, err := writer.Write([]byte(name))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(name) {
			fmt.Println("failed to write data")
			return
		}
	}

	fmt.Println(writer.String())
	//STOPMAIN1, OMIT
}
