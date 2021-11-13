package actions

import (
	"fmt"
	"github.com/gobuffalo/pop/v5"
	"github.com/hguerra/buffalo-clean-architecture-template/domain/user/model"
	repo "github.com/hguerra/buffalo-clean-architecture-template/infrastructure/database/repository/user"
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	user := &model.User{
		FirstName: "Heitor",
		LastName:  "Carneiro",
	}

	// Validate the data from the html form
	//err := database.DB.Create(model)
	//if err != nil {
	//	return err
	//}

	// OR OPEN IN VIEW --> app.go

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Validate the data from the html form
	//err := tx.Create(user)
	//if err != nil {
	//	return err
	//}

	err := repo.Db.Create(tx, user)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, r.JSON(user))
}
