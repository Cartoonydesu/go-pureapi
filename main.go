package main

import (
	"log"
	"net/http"
)

// func main() {
// 	db := database.NewDB()
// 	defer db.Close()
// 	h := skill.NewHandler(db)
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("it come here")
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`<h1>Cartoon</h1>`))
		} else if r.Method == "POST" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message": "This is POST request"}`))
		}else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"name": "anuchit"}`))
		}
	})
	log.Println("start server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
