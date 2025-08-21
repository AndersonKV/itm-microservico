package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() *sqlx.DB {
	var err error

	// Ajuste conforme seu ambiente
	connStr := "host=localhost port=5432 user=postgres password=fma dbname=postgres sslmode=disable"

	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco: %v", err)
	}

	// Testa se a conexão realmente funciona
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Erro ao fazer ping no banco: %v", err)
	}

	fmt.Println("✅ Conectado ao Postgres com sucesso!")

	// Cria tabela se não existir
	createUserTable()

	return DB
}

func createUserTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		username VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		profile_pic TEXT,
		description TEXT,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Erro ao criar tabela users: %v", err)
	}

	fmt.Println("✅ Tabela users criada/verificada com sucesso!")
}
