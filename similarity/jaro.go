package similarity

import "unicode/utf8"

type Jaro struct{}

func (j *Jaro) CompareAscii(s1, s2 string) float64 {
}

func (j *Jaro) CompareUtf8(s1, s2 string) float64 {
	mw := (utf8.RuneCountInString(s1)+utf8.RuneCountInString(s2))/2 - 1
}
