package native

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/dinhphu28/dictionary"
	"github.com/dinhphu28/dictionary/native"
)

func RunNative() {
	// 🔒 CRITICAL: never write logs to stdout
	log.SetOutput(os.Stderr)
	log.Println("Native host started")

	dictionary.StartEngine()
	// approximateLookup := dictionary.GetApproximateLookup()

	ready := dictionary.Ready()
	loadedDictionaries := dictionary.LoadedDictionaries()

	for {
		raw, err := native.ReadMessage()
		if err != nil {
			if err == io.EOF {
				log.Println("Chrome disconnected, exiting")
				return
			}
			log.Printf("read error: %v", err)
			return
		}

		var req native.Request
		if err := json.Unmarshal(raw, &req); err != nil {
			log.Printf("bad request: %v", err)
			_ = native.WriteMessage(native.Response{
				Type:    native.Error,
				Message: "invalid request",
			})
			continue
		}

		log.Printf("received: %+v", req)

		switch req.Type {

		case native.Ping:
			_ = native.WriteMessage(native.Response{
				Type:    native.Pong,
				Ready:   ready,
				Message: "Dictionaries loaded: " + strconv.Itoa(loadedDictionaries),
			})

		case native.Lookup:
			// 🔁 TEMP: fake result to prove Chrome works
			// result, err := approximateLookup.LookupWithSuggestion(req.Query)
			result, err := dictionary.Lookup(req.Query)
			if err != nil {
				_ = native.WriteMessage(native.Response{
					Type:    native.Error,
					Message: "lookup error: " + err.Error(),
				})
				continue
			}
			_ = native.WriteMessage(native.Response{
				Type:   native.Result,
				Ready:  true,
				Query:  req.Query,
				Result: result,
			})

		default:
			_ = native.WriteMessage(native.Response{
				Type:    native.Error,
				Message: "unknown message type",
			})
		}
	}
}
