package strsim

import (
	"github.com/antlabs/strsim/similarity"
)

func Hamming() OptionFunc {
	return OptionFunc(func(o *option) {

		d := &similarity.Hamming{}
		o.cmp = d.CompareUtf8
		if o.ascii {
			o.cmp = d.CompareAscii
		}
	})
}
