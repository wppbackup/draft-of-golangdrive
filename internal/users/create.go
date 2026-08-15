package users

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

func insert(db *sql.DB, u *User) (id int64, err error) {

	stmt := `INSERT INTO users ("name", "login", "password", "modified_at") VALUES ($1, $2, $3, $4) RETURNING id`

	err = db.QueryRow(stmt, u.Name, u.Login, u.Password, u.ModifiedAt).
		Scan(&id)
	if err != nil {
		return -1, err
	}

	return id, nil

}

func (h *handler) Create(rw http.ResponseWriter, r *http.Request) {

	u := new(User)

	/* r.Body is a io.Reader. For example, a *strings.Reader (for data such as plain
	 * text), *os.File (for upladed files), and so on */
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.SetHashedOriginalPassword(u.Password)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	err = u.Validate()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	u.ModifiedAt = time.Now()

	id, err := insert(h.db, u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	u.ID = id

	response := UserCreationrResponse{
		ID:         u.ID,
		Name:       u.Name,
		Login:      u.Login,
		CreatedAt:  u.CreatedAt,
		ModifiedAt: u.ModifiedAt,
		LastLogin:  u.LastLogin,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(rw).Encode(response)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

}
