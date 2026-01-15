package main

import (
	"log"
	"os"

	_ "modernc.org/sqlite"

	"github.com/dinhphu28/dictionary/api"
	"github.com/dinhphu28/dictionary/doctor"
	"github.com/dinhphu28/dictionary/native"
)

func main() {
	log.SetOutput(os.Stderr)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "doctor":
			doctor.RunDoctor()
			return

		case "native":
			native.RunNative()
			return

		case "http":
			api.RunHTTP()
			return
		}
	}

	// Default behavior
	native.RunNative()
}
