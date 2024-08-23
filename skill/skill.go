package skill

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/lib/pq"
)

type Handler struct {
	Db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{Db: db}
}

type Skill struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Tags        []string `json:"tags"`
}

func (h *Handler) GetAllSkills(w http.ResponseWriter, r *http.Request) {
	rows, err := h.Db.Query("SELECT key, name, description, logo, tags FROM skill;")
	if err != nil {
		log.Panic(err)
	}
	var skills []Skill
	for rows.Next() {
		var s Skill
		if err := rows.Scan(&s.Key, &s.Name, &s.Description, &s.Logo, pq.Array(&s.Tags)); err != nil {
			log.Panic(err)
		}
		skills = append(skills, s)
	}
	j, _ := json.Marshal(skills)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func (h *Handler) GetSkillById(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	key, _ := strings.CutPrefix(url, "/api/v1/skills/")
	if key == "" {
		log.Panic("no key found")
	}
	row := h.Db.QueryRow(fmt.Sprintf("SELECT key, name, description, logo, tags FROM skill WHERE key = '%v';", key))
	var s Skill
	if err := row.Scan(&s.Key, &s.Name, &s.Description, &s.Logo, pq.Array(&s.Tags)); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Skill not existed"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(s)
	w.Write(j)
}

func (h *Handler) CreateSkill(w http.ResponseWriter, r *http.Request) {
	var s Skill
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		log.Panic("cant extract json")
		return
	}
	stmt, err := h.Db.Prepare("INSERT INTO skill (key, name, description, logo, tags) VALUES ($1, $2, $3, $4, $5);")
	if err != nil {
		log.Panic(err)
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(s.Key, s.Name, s.Description, s.Logo, pq.Array(s.Tags)); err != nil {
		log.Panic(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	j, _ := json.Marshal(s)
	w.Write(j)
}

type UpdateSkill struct {
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Logo        string   `json:"logo" binding:"required"`
	Tags        []string `json:"tags" binding:"required"`
}

func (h *Handler) UpdateSkill(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateSkillName(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateSkillDescription(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateSkillLogo(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateSkillTags(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteSkill(w http.ResponseWriter, r *http.Request) {

}
