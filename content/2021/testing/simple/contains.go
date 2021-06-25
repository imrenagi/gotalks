package simple_test

//START, OMIT
import (
	"testing" //OMIT

	"github.com/stretchr/testify/assert" // HL
)

func TestGetData(t *testing.T) {
	data := GetData()
	assert.Len(t, data, 1)       // HL
	assert.Contains(t, data, 10) // HL
}

//STOP, OMIT

//STARTCONTAINS10, OMIT
func TestGetData_Contains10(t *testing.T) {
	data := GetData()
	var found bool
	for _, d := range data { // HL
		if d == 10 {
			found = true
		}
	}

	if !found { // HL
		t.Fail()
	}
}

//STOPCONTAINS10, OMIT

// STARTMAP, OMIT
func TestGetMap_FindTwoValues(t *testing.T) {
	m := GetMap()
	vA, ok := m["a"] // HL
	if !ok {
		t.Fail()
	}
	if vA != "valueOfA" { // HL
		t.Fail()
	}

	vB, ok := m["b"] // HL
	if !ok {
		t.Fail()
	}
	if vB != "valueOfB" { // HL
		t.Fail()
	}
}

// STOPMAP, OMIT

// STARTMAP2, OMIT
func TestGetMap_FindValue(t *testing.T) {
	m := GetMap()
	assert.Contains(t, m, "a")
	assert.Equal(t, "valueOfA", m["a"])
	assert.Contains(t, m, "b")
	assert.Equal(t, "valueOfB", m["b"])
}

// STOPMAP2, OMIT

func GetData() []int {
	return []int{}
}

func GetMap() map[string]string {
	return make(map[string]string)
}
