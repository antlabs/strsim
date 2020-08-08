package strsim

// ngram 是筛子系数需要用的一个值
func DiceCoefficient(ngram ...int) OptionFunc {
	return OptionFunc(func(o *option) {
		if len(ngram) > 0 {
			o.ngram = ngram[0]
		}
	})
}
