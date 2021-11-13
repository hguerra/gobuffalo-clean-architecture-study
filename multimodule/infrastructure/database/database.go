package database

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop/v5"
)

// DB is a connection to your database to be used
// throughout your application.
//
// https://gobuffalo.io/en/docs/db/relations#existing-model
// https://github.com/zenseth/pop-sqlite-test/blob/master/models/test_test.go
// https://github.com/gobuffalo/buffalo-pop
// https://github.com/trackq/httprouter-do/blob/main/main.go
// https://github.com/gobuffalo/pop/issues/139
var DB *pop.Connection

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")

	err = pop.AddLookupPaths("../../infrastructure/database")
	if err != nil {
		log.Fatal(err)
	}

	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}

	pop.Debug = env == "development"
}

