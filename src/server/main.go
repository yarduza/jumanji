package main

import (
	"Main-Project-1-Go/src/global"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(HandleCORS)

	if err := http.ListenAndServe(global.CurrentPort, router); err != nil {
		panic(err)
	}
}

func HandleCORS(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Access-Control-Request-Method") != "" {
		// Set CORS headers

		header := w.Header()
		header.Set("Access-Control-Allow-Credentials", "true")
		header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Anonymous")
		header.Set("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT")
		header.Set("Access-Control-Allow-Origin", global.CORSOrigin)
		header.Set("Access-Control-Max-Age", "3600")
	}

	// Adjust status code to 204
	w.WriteHeader(http.StatusNoContent)
}
