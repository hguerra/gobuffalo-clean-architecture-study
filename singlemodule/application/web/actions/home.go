package actions

import (
	"fmt"
	"github.com/gobuffalo/pop/v5"
	"net/http"
	"singlemodule/domain/user"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	model := &user.User{}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	err := tx.Create(model)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, r.JSON(model))
}
