package doctor

import (
	"fmt"

	"github.com/dinhphu28/dictionary"
)

func checkLookup() {
	dictionary.StartEngine()
	result, err := dictionary.Lookup("hello")
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
