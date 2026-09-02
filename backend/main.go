package main

import (
	"log"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func main() {
	fs := http.FileServer(http.Dir("./media"))

	http.Handle("/media/", http.StripPrefix("/media/", fs))

	log.Println("[!] http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", enableCORS(http.DefaultServeMux)))
}
