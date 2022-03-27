package db

import (
	"database/sql"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/counterposition/learngo/ent"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func NewEntClient(uri string) *ent.Client {
	database, err := sql.Open("pgx", uri)
	if err != nil {
		log.Panic(err)
	}

	driver := entsql.OpenDB(dialect.Postgres, database)
	return ent.NewClient(ent.Driver(driver))
}
