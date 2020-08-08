package similarity

type diceCoefficient struct{}

type value struct {
	shortCount int
	longCount  int
}

func (d *diceCoefficient) CompareAscii(s1, s2 string) float64 {
	return d.CompareUtf8(s1, s2)
}

func (d *diceCoefficient) CompareUtf8(s1, s2 string) float64 {
	if s1 == s2 {
		return 1.0
	}

	set1 := make(map[rune]value, len(s1)/3)

	short := s1
	long := s2
	if len(s2) < len(s1) {
		short = s2
		long = s1
	}

	l1 := 0
	for _, code := range short {
		l1++
		val, ok := set1[code]
		if !ok {
			set1[code] = value{}
		}
		val.shortCount++
		set1[code] = val
	}

	mixed := 0
	l2 := 0
	for _, code := range long {
		l2++
		val, ok := set1[code]
		if !ok {
			continue
		}

		val.longCount++
		if val.shortCount >= val.longCount {
			mixed++
		}
		set1[code] = val
	}

	return 2.0 * float64(mixed) / float64(l1+l2)
}
