package config

import (
	"context" // controla o ciclo de vida das operaçãoes com o banco
	"log"
	"os"
	"path/filepath"
	"sort" // ordenação dos arquivos de migration

	"github.com/jackc/pgx/v5/pgxpool" // pool de conexões com o postgres
)

func RunMigrations(db *pgxpool.Pool) {
	//lê todos os arquivos .sql da pasta db/migrations, ordena, e executa um por um no postgres
	files, err := os.ReadDir("db/migrations") // entra na pasta, retorna uma lista de arquivos
	if err != nil {
		log.Fatal(err)
	}

	// ordena os arquivoa pelo nome
	// garante que as migrations sejam executadas na ordem correta
	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	// percorre cada arquivo da pasta migration
	for _, file := range files {
		// monta o caminho completo do arquivo SQL -> ex: db/migrations/001_create_users.sql
		path := filepath.Join("db/migrations", file.Name())
		// lê todo o conteúdo do arquivo SQL para a memória
		sql, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		// executa o SQL no postgres
		// context.Background() é um contexto básico, sem timeout
		_, err = db.Exec(context.Background(), string(sql))
		if err != nil {
			log.Fatal("erro na migration %s: %v", file.Name(), err)
		}
		log.Println("Migration executada:", file.Name())
	}
}
