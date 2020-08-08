package strsim

import (
	"github.com/antlabs/strsim/similarity"
)

// 比较两个字符串相似度
func Compare(s1, s2 string, opts ...Option) float64 {
	var o option

	o.fillOption()

	return compare(s1, s2, &o)
}

// 返回相似度最高的那个字符串
func FindBestMatch(s string, targets []string, opts ...Option) *similarity.MatchResult {
	return findBestMatch(s, targets, opts...)
}

// 比较两个字符串内部函数
func compare(s1, s2 string, o *option) float64 {
	if s, e := modifyStrAndCheck(o, &s1, &s2); e {
		return s
	}

	return o.cmp(s1, s2)
}

// 前处理主要涉及，修改字符串，和边界判断
func modifyStrAndCheck(o *option, s1, s2 *string) (score float64, exit bool) {
	modifyString(o, s1)
	modifyString(o, s2)

	return check(*s1, *s2)
}

// 返回相似度最高的那个字符串，内部函数
func findBestMatch(s string, targets []string, opts ...Option) *similarity.MatchResult {

	var opt option
	opt.fillOption()

	match := make([]*similarity.Match, 0, len(targets))
	bestIndex := 0
	for k, s2 := range targets {

		score := compare(s, s2, &opt)

		match = append(match, &similarity.Match{S: s2, Score: score})

		if k == 0 {
			continue
		}

		if score > match[bestIndex].Score {
			bestIndex = k
		}
	}

	return &similarity.MatchResult{AllResult: match, Match: match[bestIndex], BestIndex: bestIndex}
}
