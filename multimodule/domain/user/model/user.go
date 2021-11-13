package model

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

// User is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}
