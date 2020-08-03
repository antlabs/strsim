package strsim

type edit struct {
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (e *edit) Compare(s1, s2 string) float64 {
	cacheX := make([]int, len(s2))

	diagonal := 0
	for y, yLen := 0, len(s1); y < yLen; y++ {
		for x, xLen := 0, len(cacheX); x < xLen; x++ {
			on := x + 1
			left := y + 1

			if y == 0 {
				diagonal = x
			}
			if y > 0 {
				on = cacheX[x]
				if x-1 >= 0 {
					left = cacheX[x-1]
				}
			}

			same := 0
			if s1[y] != s2[x] {
				same = 1
			}

			cacheX[x] = min(min(on, left), same+diagonal)
			diagonal = cacheX[x]
		}
	}

	return float64(cacheX[len(cacheX)-1]) //TODO修改
}
