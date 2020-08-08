package similarity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DiceCoefficient_CompareAscii(t *testing.T) {
	d := &diceCoefficient{}

	for k, v := range []testOneCase{
		{s1: "ivan1", s2: "ivan2", cost: 0.8},
		{s1: "love", s2: "love", cost: 1},
	} {
		assert.Equal(t, d.CompareAscii(v.s1.(string), v.s2.(string)), v.cost, fmt.Sprintf("error case:%d", k))
	}
}

func Test_DiceCoefficient_CompareUtf8(t *testing.T) {
	d := &diceCoefficient{}

	for k, v := range []testOneCase{
		{s1: "你好中国", s2: "你好中国", cost: 1},
		{s1: "加油，来个", s2: "加油，来", cost: 0.8},
	} {
		assert.Equal(t, d.CompareUtf8(v.s1.(string), v.s2.(string)), v.cost, fmt.Sprintf("error case:%d", k))
	}
}

func Test_DiceCoefficient_FindBestMatch(t *testing.T) {
	d := &diceCoefficient{}

	for k, v := range []testBestCase{
		{s: "白日依山尽", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 0},
		{s: "黄河流", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 1},
		{s: "一层", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 3},
		{s: "楼", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 3},
		{s: "山近", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 0},
		{s: "海刘", targets: []string{"白日依山尽", "黄河入海流", "欲穷千里目", "更上一层楼"}, bestIndex: 1},
	} {
		mr := findBestMatch(v.s, v.targets, d.CompareUtf8)
		assert.Equal(t, mr.BestIndex, v.bestIndex, fmt.Sprintf("error case:%d", k))
	}
}
