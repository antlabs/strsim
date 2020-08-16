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
		j.MatchWindow = 0
	}

}

func Test_Jaro_CompareUtf8(t *testing.T) {
	j := &Jaro{}

	// jaro处理两个字符串长串接近的数据会好点
	for k, v := range []testOneCase{

		{s1: "二一三四五", s2: "一二五四", cost: 0.6722222222222222},
		{s2: "二一三四五", s1: "一二五四", cost: 0.6722222222222222},
		{s1: "中文也被称为华文、汉文。中文（汉语）有标准语和方言之分，其标准语即汉语普通话", s2: "中文", cost: 0.6842105263157894, match: 1},
	} {
		m := fmt.Sprintf("error case:%d", k)
		assert.Equal(t, j.CompareUtf8(v.s1, v.s2), v.cost, m)
		j.mw = 0
	}

}
