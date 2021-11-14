package actions

import (
	"fmt"
	"github.com/gobuffalo/logger"
	"github.com/gobuffalo/pop/v5"
	"net/http"
	"singlemodule/domain/user"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	// https://github.com/gobuffalo/buffalo/blob/28bd05f88648468281d76c439c36f8ff2482cf91/default_context_test.go#L20
	// https://github.com/gobuffalo/buffalo/blob/28bd05f88648468281d76c439c36f8ff2482cf91/options.go#L135
	log := logger.New(logger.DebugLevel).WithField("handler", "HomeHandler")
	log.Info("Running handler...")

	log = c.Logger().WithField("handler", "HomeHandler")
	log.Info("Running handler with request_id...")

	c.Logger().Info("Sava user in database")
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

	c.Logger().Info("User user saved successfully")
	return c.Render(http.StatusOK, r.JSON(model))
}
