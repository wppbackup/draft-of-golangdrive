package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func insert(db *sql.DB, u *User) (int64, error) {

	stmt := `insert into ("name", "logn", "password", "modified_at") VALUES ($1, $2, $3, $4)`
	result, err := db.Exec(stmt, u.Name, u.Login, u.Password, u.ModifiedAt)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()

}

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {

	u := new(User)
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	u.SetEncryptedPassword(u.Password)

	err = u.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	id, err := insert(h.db, u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
	u.ID = id

	rw.Header().Add("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(u)
}
