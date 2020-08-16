package similarity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testOneCase struct {
	s1    string
	s2    string
	cost  float64
	ngram int
	mixed int
	match int
}

type testBestCase struct {
	s         string
	targets   []string
	bestIndex int
}

func Test_EditDistance_CompareAscii(t *testing.T) {
	e := &EditDistance{}

	for k, v := range []testOneCase{
		{s1: "ivan1", s2: "ivan2", cost: 0.8, mixed: 1},
		{s1: "love", s2: "love", cost: 1, mixed: 0},
		{s1: "kitten", s2: "sitting", cost: 1 - 3/7.0, mixed: 3},
		{s1: "12", s2: "1", cost: 0.5, mixed: 1},
		{s1: "1", s2: "12", cost: 0.5, mixed: 1},
		{s1: "123", s2: "1", cost: 0.33333333333333337, mixed: 2},
		{s1: "1", s2: "123", cost: 0.33333333333333337, mixed: 2},
		{s1: "1234", s2: "1", cost: 0.25, mixed: 3},
		{s1: "1", s2: "1234", cost: 0.25, mixed: 3},
	} {
		s := e.CompareAscii(v.s1, v.s2)
		assert.Equal(t, s, v.cost, fmt.Sprintf("cost:error case:%d", k))
		assert.Equal(t, e.mixed, v.mixed, fmt.Sprintf("mixed:error case:%d", k))
	}
}

func Test_EditDistance_CompareUtf8(t *testing.T) {
	e := &EditDistance{}

	for k, v := range []testOneCase{
		{s1: "你好中国", s2: "你好中国", cost: 1, mixed: 0},
		{s1: "加油，来个", s2: "加油，来", cost: 0.8, mixed: 1},
		{s1: "一二三三四五", s2: "六二三三二五七", cost: 1 - 3/7.0, mixed: 3},
		{s1: "一二", s2: "一", cost: 0.5, mixed: 1},
		{s1: "一", s2: "一二", cost: 0.5, mixed: 1},
		{s1: "一二三", s2: "一", cost: 0.33333333333333337, mixed: 2},
		{s1: "一", s2: "一二三", cost: 0.33333333333333337, mixed: 2},
		{s1: "一二三四", s2: "一", cost: 0.25, mixed: 3},
		{s1: "一", s2: "一二三四", cost: 0.25, mixed: 3},
		{s1: "中文也被称为华文、汉文。中文（汉语）有标准语和方言之分，其标准语即汉语普通话", s2: "方块", cost: 0.02631578947368418, mixed: 37},
	} {
		s := e.CompareUtf8(v.s1, v.s2)
		assert.Equal(t, s, v.cost, fmt.Sprintf("cost:error case:%d", k))
		assert.Equal(t, e.mixed, v.mixed, fmt.Sprintf("mixed:error case:%d", k))
	}
}

func Test_EditDistance_FindBestMatch(t *testing.T) {
	e := &EditDistance{}

	for k, v := range []testBestCase{
		{s: "白日依山尽", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 0},
		{s: "黄河流", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 1},
		{s: "一层", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 3},
		{s: "楼", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 3},
		{s: "山近", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 0},
		{s: "海刘", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 1},
	} {
		mr := findBestMatch(v.s, v.targets, e.CompareUtf8)
		assert.Equal(t, mr.BestIndex, v.bestIndex, fmt.Sprintf("error case:%d", k))
	}
}
