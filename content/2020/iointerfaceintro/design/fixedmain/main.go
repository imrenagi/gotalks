package main

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/imrenagi/gotalks/content/2020/iointerfaceintro/design/blob"
)

func main() {
	// STARTMAINFIXED, OMIT
	p1 := Person{"Imre"}
	f1, _ := ioutil.TempFile("", "")
	defer f1.Close()

	err := p1.Write(f1)
	// handling error
	// STOPMAINFIXED, OMIT

	// STARTMAINGCS, OMIT
	p2 := Person{"Nagi"}
	gcs, _ := blob.NewGCS("bucket", "person.txt")
	err := p2.Write(gcs.Writer())
	// handling error
	// STOPMAINGCS, OMIT
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

//STARTPERSONWRITE, OMIT
func (p Person) Write(w io.WriteCloser) error {
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

//STOPPERSONWRITE, OMIT
