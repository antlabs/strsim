package similarity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hamming_CompareAscii(t *testing.T) {
	h := Hamming{}

	for k, v := range []testOneCase{
		{s1: "1011101000", s2: "1001001000", cost: 0.8},
		{s1: "21438960", s2: "22337960", cost: 0.625},
		{s1: "toned", s2: "roses", cost: 0.4},
	} {
		assert.Equal(t, h.CompareAscii(v.s1, v.s2), v.cost, fmt.Sprintf("error case:%d", k))
	}
}

func Test_Hamming_CompareUtf8(t *testing.T) {
	h := Hamming{}

	for k, v := range []testOneCase{
		{s1: "中国嘿嘿", s2: "中国哈哈", cost: 0.5},
		{s1: "中国嘿嘿1", s2: "中国哈哈", cost: 0.4},
	} {
		assert.Equal(t, h.CompareUtf8(v.s1, v.s2), v.cost, fmt.Sprintf("error case:%d", k))
	}
}
