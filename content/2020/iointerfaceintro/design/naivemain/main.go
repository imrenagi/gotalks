package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//STARTMAIN, OMIT
	p1 := Person{"Imre"}

	b, _ := json.Marshal(p1)
	f, _ := ioutil.TempFile("", "")
	defer f.Close()
	fmt.Println(f.Name())

	n, err := f.Write(b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "bytes written")
	//STOPMAIN, OMIT
}

func incorrectsample() {
	//STARTINEFFECTIVESAMPLE, OMIT
	p1 := Person{"Imre"}
	b, _ := json.Marshal(p1)

	f1, _ := ioutil.TempFile("", "")
	f2, _ := os.Create("testing.txt")

	n1, err := f1.Write(b)
	// handling error

	n2, err := f2.Write(b)
	// handling error
	//STOPINEFFECTIVESAMPLE, OMIT
}

type People []Person

func (p People) Write(w io.WriteCloser) error {
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		return err
	}
	return w.Close()
}

//STARTPERSON, OMIT
type Person struct {
	Name string `json:"name"`
}

//STOPPERSON, OMIT

func (p Person) Write(w io.Writer) error {
	b, err := json.Marshal(p)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		return err
	}

	return nil
}
