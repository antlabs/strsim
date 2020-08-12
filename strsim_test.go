package strsim

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Compare_Special(t *testing.T) {

	for _, v := range []testCase{
		{arg1: "", arg2: "", got: 1},
		{arg1: "", arg2: "", got: 1, opt: Hamming()},
		{arg1: "", arg2: "", got: 1, opt: Jaro()},
		{arg1: "", arg2: "", got: 1, opt: DiceCoefficient()},

		{arg1: "1", arg2: "", got: 0},
		{arg1: "1", arg2: "", got: 0, opt: Hamming()},
		{arg1: "1", arg2: "", got: 0, opt: Jaro()},
		{arg1: "1", arg2: "", got: 0, opt: DiceCoefficient()},

		{arg1: "", arg2: "1", got: 0},
		{arg1: "", arg2: "1", got: 0, opt: Hamming()},
		{arg1: "", arg2: "1", got: 0, opt: Jaro()},
		{arg1: "", arg2: "1", got: 0, opt: DiceCoefficient()},
	} {
		sim := Compare(v.arg1, v.arg2)
		assert.Equal(t, v.got, sim)
	}
}
