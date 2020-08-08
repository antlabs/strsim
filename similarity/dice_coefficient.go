package similarity

import (
	"strings"
	"unicode/utf8"
)

type DiceCoefficient struct {
	Ngram int
}

type value struct {
	s1Count int
	s2Count int
}

func (d *DiceCoefficient) CompareAscii(s1, s2 string) float64 {
	return d.CompareUtf8(s1, s2)
}

func (d *DiceCoefficient) setOr(set map[string]value, s string, s1 bool) (mixed, l int) {
	var key strings.Builder
	ngram := d.Ngram
	if ngram == 0 {
		ngram = 1
	}

	for i := 0; i < len(s); {
		currSize := 0
		for j, total := 0, 0; j < ngram; j++ {
			r, size := utf8.DecodeRuneInString(s[i+total:])
			key.WriteRune(r)
			total += size
			if j == 0 {
				currSize = size
			}

			if i+total >= len(s) {
				break
			}

			l++

			val, ok := set[key.String()]
			if s1 {
				if !ok {
					val = value{}
				}
				val.s1Count++
			} else {

				if !ok {
					continue
				}

				val.s2Count++
				if val.s1Count >= val.s2Count {
					mixed++
				}
			}

			set[key.String()] = val

			key.Reset()
		}

		i += currSize
	}

	return mixed, l
}

func (d *DiceCoefficient) CompareUtf8(s1, s2 string) float64 {

	set := make(map[string]value, len(s1)/3)
	//TODO 边界比如字符长度小于ngram

	mixed, l1 := d.setOr(set, s1, true)

	mixed, l2 := d.setOr(set, s2, false)

	return 2.0 * float64(mixed) / float64(l1+l2)
}
