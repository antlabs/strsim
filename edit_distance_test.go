package strsim

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	s1   interface{}
	s2   interface{}
	cost float64
}

func Test_EditDistance_Compare(t *testing.T) {
	e := &edit{}

	for k, v := range []testCase{
		{s1: "ivan1", s2: "ivan2", cost: 0.8},
		{s1: "love", s2: "love", cost: 1},
	} {
		assert.Equal(t, e.CompareAscii(v.s1.(string), v.s2.(string)), v.cost, fmt.Sprintf("error case:%d", k))
	}
}

func Test_EditDistance_CompareRune(t *testing.T) {
	e := &edit{}

	for k, v := range []testCase{
		{s1: "你好中国", s2: "你好中国", cost: 1},
		{s1: "加油，来个", s2: "加油，来", cost: 0.8},
	} {
		assert.Equal(t, e.CompareUtf8(v.s1.(string), v.s2.(string)), v.cost, fmt.Sprintf("error case:%d", k))
	}
}
