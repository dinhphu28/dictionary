package api

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	lookupHandlerV2 LookupHandlerV2
}

func NewRouter(lookupHandlerV2 LookupHandlerV2) *Router {
	return &Router{
		lookupHandlerV2: lookupHandlerV2,
	}
}

func (r *Router) StartAPIServer() {
	mux := http.NewServeMux()
	handler := CorsMiddleware(mux)

	mux.Handle("/v2/lookup", &r.lookupHandlerV2)

	fmt.Println("Listening at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
