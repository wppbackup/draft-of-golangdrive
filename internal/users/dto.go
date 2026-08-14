package users

import (
	"time"
)

type UserCreationrResponse struct {
    ID         int64     `json:"id"`
    Name       string    `json:"name"`
    Login      string    `json:"login"`
    CreatedAt  time.Time `json:"created_at"`
    ModifiedAt time.Time `json:"modified_at"`
    LastLogin  time.Time `json:"last_login"`
}