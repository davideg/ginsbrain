package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)


func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteAddr := r.RemoteAddr
		if len(remoteAddr) == 0 {
			remoteAddr = r.Header.Get("x-forwarded-for")
		}
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
  r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler).Methods("GET")

  port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	http.Handle("/", r)


	log.Printf("[+] Uh, hi! My brain is running on port %s", port)
	http.ListenAndServe(":"+port, Log(http.DefaultServeMux))
}
