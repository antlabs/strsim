package strsim

type option struct {
	ignore int
	cmp    func(s1, s2 string) float64
}

func (o *option) fillOption(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(o)
	}
}

type Option interface {
	Apply(*option)
}

type OptionFunc func(*option)

func (o OptionFunc) Apply(opt *option) {
	o(opt)
}

//忽略大小写
func IgnoreCase() OptionFunc {
	return OptionFunc(func(o *option) {
		o.ignore |= ignoreCase
	})
}

//忽略空白字符
func IgnoreSpace() OptionFunc {
	return OptionFunc(func(o *option) {
		o.ignore |= ignoreSpace
	})
}

//使用编辑距离
func Levdist() OptionFunc {
	return OptionFunc(func(o *option) {
	})
}
