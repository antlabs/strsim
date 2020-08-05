package strsim

type Match struct {
	S     string
	Score float64
}

type MatchResult struct {
	AllResult []*Match
	Match     *Match
	BestIndex int
}

type Compare func(s1, s2 string) float64
