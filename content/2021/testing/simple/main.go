//START, OMIT
package random_test // Avoid testing code to be bundled in prod

import (
	"testing"
)

func TestRandom(t *testing.T) { // Test... prefix is mandatory
	if 2 != 3 {
		t.Fail()
		t.Log("2 is not equal to 3")
	}
}

//STOP, OMIT
