package strsim

import "github.com/antlabs/strsim/similarity"

// CosineConf is a configuration struct for Cosine similarity.

func Cosine() OptionFunc {
	return OptionFunc(func(o *option) {
		h := &similarity.Cosine{}
		o.base64 = true
		o.cmp = h.CompareUtf8
		if o.ascii {
			o.cmp = h.CompareAscii
		}
	})

}
