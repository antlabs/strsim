package similarity

import (
	"fmt"
	"math"
	"unicode/utf8"
)

type Jaro struct {
	// test use
	mw int
	m  int
	t  int
}

func (j *Jaro) CompareAscii(s1, s2 string) float64 {
	return j.CompareUtf8(s1, s2)
}

func (j *Jaro) CompareUtf8(s1, s2 string) float64 {
	mw := max(utf8.RuneCountInString(s1), utf8.RuneCountInString(s2))/2 - 1
	m := 0

	set := make(map[rune][]int, len(s1)/3)
	l1 := 0
	for k, c := range s1 {
		set[c] = append(set[c], k)
		l1++
	}

	t := 0
	l2 := 0

	best := 0
	for s2Index, c := range s2 {
		indexs, ok := set[c]
		l2++
		if !ok {
			continue
		}

		for _, i := range indexs {
			if math.Abs(float64(s2Index-i)) <= float64(mw) {
				m++
				if i < best {
					t++
					fmt.Printf("--->%c:s2Index(%d):i(%d), s2Index-i(%d), mw:%d\n", c, s2Index, i, s2Index-i, mw)
				}
				best = i

				break
			}
		}
	}

	m2 := float64(m)
	j.mw = mw
	j.m = m
	j.t = t
	return 1.0 / 3.0 * (m2/float64(l1) + m2/float64(l2) + (m2-float64(t)/2.0)/m2)
}
