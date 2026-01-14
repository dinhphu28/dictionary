package ranking

func Levenshtein(a, b string) int {
	ar := []rune(a)
	br := []rune(b)

	da := make([][]int, len(ar)+1)
	for i := range da {
		da[i] = make([]int, len(br)+1)
	}

	for i := 0; i <= len(ar); i++ {
		da[i][0] = i
	}
	for j := 0; j <= len(br); j++ {
		da[0][j] = j
	}

	for i := 1; i <= len(ar); i++ {
		for j := 1; j <= len(br); j++ {
			cost := 0
			if ar[i-1] != br[j-1] {
				cost = 1
			}

			da[i][j] = min3(
				da[i-1][j]+1,      // deletion
				da[i][j-1]+1,      // insertion
				da[i-1][j-1]+cost, // substitution
			)
		}
	}
	return da[len(ar)][len(br)]
}

func min3(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= c {
		return b
	}
	return c
}
