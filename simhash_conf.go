package strsim

import "github.com/antlabs/strsim/similarity"

func Simhash() OptionFunc {
	return OptionFunc(func(o *option) {
		h := &similarity.Simhash{}
		o.base64 = true
		o.cmp = h.CompareUtf8
		if o.ascii {
			o.cmp = h.CompareAscii
		}
	})

}
