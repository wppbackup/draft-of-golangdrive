package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

func insert(db *sql.DB, u *User) (int64, error) {
	
	stmt := `INSERT INTO users ("name", "login", "password", "modified_at") VALUES ($1, $2, $3, $4) RETURNING id`
	
	var id int64
	err := db.QueryRow(stmt, u.Name, u.Login, u.Password, u.ModifiedAt).Scan(&id)
	if err != nil {
		return -1, err
	}
	
	return id, nil
}

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {

	u := new(User)
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.SetEncryptedPassword(u.Password)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	u.ModifiedAt = time.Now()
	
	err = u.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := insert(h.db, u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	u.ID = id

	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader((http.StatusCreated))
	json.NewEncoder(rw).Encode(u)
}
