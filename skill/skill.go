package skill

import (
	"database/sql"
	"net/http"
)

type Handler struct {
	Db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{Db: db}
}

func (h *Handler) GetAllSkill(w http.ResponseWriter, req *http.Request) {

}