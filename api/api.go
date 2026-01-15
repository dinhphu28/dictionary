package api

import (
	"fmt"

	"github.com/dinhphu28/dictionary"
)

func RunHTTP() {
	fmt.Println("HTTP mode")

	dictionary.StartEngine()
	lookupHandlerV2 := NewLookupHandlerV2()
	router := NewRouter(*lookupHandlerV2)
	router.StartAPIServer()
}
