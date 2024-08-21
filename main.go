package main

import (
	"cartoonydesu/database"
	"cartoonydesu/skill"
	"log"
	"net/http"
)

// func main() {
// 	db := database.NewDB()
// 	defer db.Close()
// 	h := skill.NewHandler(db)
// }

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("it come here")
// 		if r.Method == "GET" {
// 			w.Header().Set("Content-Type", "text/plain")
// 			w.WriteHeader(http.StatusOK)
// 			w.Write([]byte(`<h1>Cartoon</h1>`))
// 		} else if r.Method == "POST" {
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusOK)
// 			w.Write([]byte(`{"message": "This is POST request"}`))
// 		}else {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte(`{"name": "anuchit"}`))
// 		}
// 	})
// 	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte(`{"message": "name"}`))
// 	})
// 	http.HandleFunc("/path", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([] byte(`{"message": "path"}`))
// 	})
// 	// log.Println("start server at port :8080")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	db := database.NewDB()
	defer db.Close()
	h := skill.NewHandler(db)
	http.HandleFunc("/api/v1/skills", func(w http.ResponseWriter, r *http.Request) {
		// log.Println(r.URL.Path)
		// if r.URL.Path == "/api/v1/skills" {
			if r.Method == "GET" {
				h.GetAllSkills(w, r)
			} else {
				methodNotAllowResponse(w)
			}
		// } else {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusNotFound)
		// 	w.Write([]byte(`{"message": "Unknow request path"}`))
		// }
	})
	http.HandleFunc("/api/v1/skills/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			// log.Print(r.URL.Path)
			h.GetSkillById(w, r)
		} else {
			methodNotAllowResponse(w)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func methodNotAllowResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(`{"message": "Method not allow"}`))
}
