package repository

import "github.com/gobuffalo/pop/v5"

type CrudRepository interface {
	Create(tx *pop.Connection, entity interface{}) error
}
