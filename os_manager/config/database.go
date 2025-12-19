package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB() *pgxpool.Pool {
	dsn := "postgres://postgres:postgres@localhost:5432/os_manager?sslmode=disable"

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Println("Banco conectado!")
	return db
}
