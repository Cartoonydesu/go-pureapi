package main

import (
	"cartoonydesu/skill"
	"database/sql"
	"net/http"
)

type router struct {
	routes map[string]map[string]http.HandlerFunc
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if handlers, ok:= r.routes[req.URL.Path]; ok {
		if handler, methodExists := handlers[req.Method]; methodExists {
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

func (r *router) addRoute(method, path string, handler http.HandlerFunc) {
	if r.routes[path] == nil {
		r.routes[path] = make(map[string]http.HandlerFunc)
	}
	r.routes[path][method] = handler
}

func NewRouter(db *sql.DB) *router {
	router := &router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}
	sh := skill.NewHandler(db)
	router.addRoute("GET", "/api/v1/skills", sh.GetAllSkill)

	return router
}