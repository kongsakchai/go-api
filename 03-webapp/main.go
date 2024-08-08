package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/render", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		b := r.FormValue("text")
		w.Write([]byte(b))
	})

	http.Handle("/", http.FileServer(http.Dir("public")))

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
