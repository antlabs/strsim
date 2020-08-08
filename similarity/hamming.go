package similarity

import (
	"unicode/utf8"
)

type hamming struct{}

func (h *hamming) CompareAscii(s1, s2 string) float64 {

	count := 0
	l := 0
	for i, j := 0, 0; i < len(s1) && j < len(s2); {

		if s1[i] != s2[j] {
			count++
		}

		i++
		j++
		l++
	}

	return 1.0 - float64(count)/float64(l)
}

func (h *hamming) CompareUtf8(utf8Str1, utf8Str2 string) float64 {
	count := 0

	l := 0
	for i, j := 0, 0; i < len(utf8Str1) && j < len(utf8Str2); {
		size := 0
		r1, size := utf8.DecodeRune(StringToBytes(utf8Str1[i:]))
		i += size

		r2, size := utf8.DecodeRune(StringToBytes(utf8Str2[j:]))
		j += size

		if r1 != r2 {
			count++
		}

		l++
	}

	return 1.0 - float64(count)/float64(l)
}
