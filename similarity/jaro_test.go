package similarity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Jaro_CompareAscii(t *testing.T) {
	j := &Jaro{}

	for k, v := range []testOneCase{

		{s1: "bacde", s2: "abed", cost: 0.6722222222222222},
		{s1: "MARTHA", s2: "MARHTA", cost: 0.9444444444444444},
		{s1: "DIXON", s2: "DICKSONX", cost: 0.7666666666666666},

		{s1: "JELLYFISH", s2: "SMELLYFISH", cost: 0.8962962962962964},
		{s2: "JELLYFISH", s1: "SMELLYFISH", cost: 0.8962962962962964},
	} {
		m := fmt.Sprintf("error case:%d", k)
		assert.Equal(t, j.CompareAscii(v.s1, v.s2), v.cost, m)
	}

}
