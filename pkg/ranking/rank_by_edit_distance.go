package ranking

import (
	"sort"
	"strings"
)

type Match struct {
	Word string
	Dist int
}

func RankByEditDistance(query string, candidates []string) []Match {
	q := strings.ToLower(query)

	var list []Match
	for _, c := range candidates {
		d := Levenshtein(q, strings.ToLower(c))
		list = append(list, Match{Word: c, Dist: d})
	}

	sort.Slice(list, func(i, j int) bool {
		if list[i].Dist == list[j].Dist {
			return list[i].Word < list[j].Word
		}
		return list[i].Dist < list[j].Dist
	})

	return list
}
