package strsim

import (
	"github.com/antlabs/strsim/similarity"
)

// ngram 是筛子系数需要用的一个值
func Jaro() OptionFunc {
	return OptionFunc(func(o *option) {
		d := &similarity.Jaro{}
		o.cmp = d.CompareUtf8
	})
}
