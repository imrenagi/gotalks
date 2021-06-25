package math_test

import (
	"testing"

	"github.com/imrenagi/gotalks/content/2021/testing/math"
	"github.com/stretchr/testify/assert"
)

//START,OMIT
func TestDiv(t *testing.T) {
	assert.Equal(t, 0, math.Div(1/2))
	assert.Equal(t, 2, math.Div(4/2))
	assert.Equal(t, -2, math.Div(-8/4))
}

//STOP,OMIT

//STARTTT, OMIT
func TestDiv_Tables(t *testing.T) {
	cases := []struct{ // Test table // HL
		a, b int
		expected int
	} {
		{ 
			{a: 1, b: 0, expected: 0}, // what if this error? 
			{a: 2, b: 1, expected: 2}, 
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected, math.Div(c.a, c.b))		
	}
}
//STOPTTT, OMIT


//STARTTT2, OMIT
func TestDiv_Tables(t *testing.T) {
	cases := []struct{ 
		name string // HL
		a, b int
		expected int
	} {
		{ 
			{name: "test division with 0", a: 1, b: 0, expected: 0}, 
			{name: "positive numbers division", a: 2, b: 2, expected: 1}, 
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) { // HL
			assert.Equal(t, c.expected, math.Div(c.a, c.b))
		})		
	}
}
//STOPTT2, OMIT


//STARTSUB, OMIT
func TestRandom(t *testing.T) {
  // prepare...
  defer func() {
		//cleanUp()
	}

  t.Run("positive cases", func(t *testing.T) { // HL
    // check positive cases
    // you might want to have test table if you want
  })

  t.Run("negative cases", func(t *testing.T) { // HL
    // check negative cases
  })
}
//STOPSUB, OMIT