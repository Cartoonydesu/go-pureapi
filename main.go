package main

import (
	"cartoonydesu/database"
	"cartoonydesu/skill"
	"log"
	"net/http"
	"strings"
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
		if r.Method == "GET" {
			h.GetAllSkills(w, r)
		} else if r.Method == "POST" {
			h.CreateSkill(w, r)
		} else if r.Method == "DELETE" {
			h.DeleteSkill(w, r)
		} else {
			methodNotAllowResponse(w)
		}
	})
	http.HandleFunc("/api/v1/skills/", func(w http.ResponseWriter, r *http.Request) {
		path, _ := strings.CutPrefix(r.URL.Path, "/api/v1/skills/")
		pathArr := strings.Split(path, "/")
		if r.Method == "GET" {
			h.GetSkillById(w, r)
		} else if r.Method == "PUT" && len(pathArr) == 1 {
			if len(pathArr) == 1 {
				h.UpdateSkill(w, r, pathArr[0])
			} else {
				methodNotAllowResponse(w)
			}
		} else if r.Method == "PATCH" && len(pathArr) == 3 && pathArr[1] == "action" {
			switch pathArr[2] {
			case "name":
				h.UpdateSkillName(w, r, pathArr[0])
			case "description":
				h.UpdateSkillDescription(w, r, pathArr[0])
			case "logo":
				h.UpdateSkillLogo(w, r, pathArr[0])
			case "tags":
				h.UpdateSkillTags(w, r, pathArr[0])
			}
		} else {
			methodNotAllowResponse(w)
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		methodNotFoundResponse(w)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func methodNotAllowResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(`{"message": "Method not allow"}`))
}

func methodNotFoundResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "Path not found"}`))
}
