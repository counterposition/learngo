package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v4"
	"log"

	_ "github.com/jackc/pgx/v4"
	"go.uber.org/fx"
)

const defaultDatabaseUri = "postgres://postgres:postgres@localhost:5432/postgres"

type Connection = *pgx.Conn

func NewDatabaseClient(lifecycle fx.Lifecycle, uri string) Connection {
	conn, err := pgx.Connect(context.Background(), uri)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}

	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return conn.Close(ctx)
		},
	})

	return conn
}

func CreateUsersTable(conn Connection) error {
	_, err := conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func InsertDummyUser(conn Connection) error {
	name := "Alice"
	email := "alice@example.org"
	_, err := conn.Exec(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var databaseUri string
	flag.StringVar(&databaseUri, "database", defaultDatabaseUri, "Database to connect to")
	flag.Parse()

	app := fx.New(
		fx.Supply(databaseUri),
		fx.Provide(NewDatabaseClient),

		fx.Invoke(CreateUsersTable),
		fx.Invoke(InsertDummyUser),
	)

	// In this case it would be better to call app.Run() but I'm just showing off

	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Unable to start application: %v\n", err)
	}

	select {
	case signal := <-app.Done():
		log.Println("Application stopped with signal ", signal)
	}
}
