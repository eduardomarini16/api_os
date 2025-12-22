package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Cria e retorna uma conexão com o banco de dados postgres
func ConnectDB() *pgxpool.Pool {
	dsn := "postgres://postgres:postgres@localhost:5432/os_manager?sslmode=disable"

	// cria um pool de conexões com o bancio
	// o pool gerencia várias conexões automaticamente
	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal(err)
	}

	// testa se o banco realmente está acessível
	// ping envia uma requisiçãpo simples para validar a conexão
	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	log.Println("Banco conectado!")
	// retorna o pool de conexões para serusado no restante da aplicação
	return db
}
