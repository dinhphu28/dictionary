package api

import (
	"net/http"
	"strings"

	"dinhphu28.com/dictionary/dictionary"
)

type LookupHandlerV2 struct {
	approximateLookup dictionary.ApproximateLookup
}

func NewLookupHandlerV2(approximateLookup dictionary.ApproximateLookup) *LookupHandlerV2 {
	return &LookupHandlerV2{
		approximateLookup: approximateLookup,
	}
}

func (lookupHandler *LookupHandlerV2) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	q := strings.TrimSpace(r.URL.Query().Get("q"))
	if q == "" {
		http.Error(w, "missing q parameter", http.StatusBadRequest)
		return
	}

	result, err := lookupHandler.approximateLookup.LookupWithSuggestion(q)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if len(result.LookupResults) == 0 {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	writeJSONResponse(w, result)
}
