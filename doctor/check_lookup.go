package doctor

import (
	"fmt"

	"dinhphu28.com/dictionary"
)

func checkLookup() {
	dictionary.StartEngine()
	approximateLookup := dictionary.GetApproximateLookup()

	result, err := approximateLookup.LookupWithSuggestion("hello")
	if err != nil {
		fmt.Println("✖ Lookup failed:", err)
		return
	}

	if len(result.LookupResults) == 0 {
		fmt.Println("⚠ Lookup returned no results")
		return
	}

	fmt.Printf(
		"✔ Lookup test passed (\"hello\" → %s)\n",
		result.LookupResults[0].Dictionary,
	)
}
