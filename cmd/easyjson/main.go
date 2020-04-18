package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"time"

	"github.com/amin-mir/jsonperfopt/person"
	"github.com/mailru/easyjson"
)

const (
	port     int = 8080
	minSleep     = 15
	maxSleep     = 50
)

func sleep() {
	dur := rand.Intn(maxSleep-minSleep+1) + minSleep
	time.Sleep(time.Duration(dur) * time.Millisecond)
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method not alloed", http.StatusMethodNotAllowed)
	}

	var p person.Person
	if err := easyjson.UnmarshalFromReader(r.Body, &p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	_, _, err := easyjson.MarshalToHTTPResponseWriter(p, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", EchoHandler)

	// Register pprof handlers
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	log.Printf("Starting HTTP server on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatal(err)
	}
}
