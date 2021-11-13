package user

import (
	"github.com/gobuffalo/pop/v5"
	"github.com/hguerra/buffalo-clean-architecture-template/domain/user/model"
)

type Repository struct{}

var Db *Repository

func init() {
	Db = &Repository{}
}

// Create persist data
func (r *Repository) Create(tx *pop.Connection, user *model.User) error {
	return tx.Create(user)
}
