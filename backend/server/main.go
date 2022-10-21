package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	m := make(map[string]Hub)

	http.HandleFunc("/create-hub", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		hub := newHub()
		go hub.run()
		m[code] = *hub
		serveWs(hub, w, r)
	})

	http.HandleFunc("/join-hub", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		jHub := m[code]
		serveWs(&jHub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
