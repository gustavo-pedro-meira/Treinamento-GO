package database

import (
	"database/sql"
	"fmt"
	// O _ importa o driver, registrando sem usar diretamente
	_ "github.com/jackc/pgx/v5/stdlib"
)

// O *sql.DB é um pool de conexões. É um obj que gerencia varias conexões com o bd
// O error, se algo der errado durante o processo, a func retonará um erro
func Conectar() (*sql.DB, error) {
	// DSN é uma stringque contem todas info p encontrar e acessar o BD
	// Formato: tipodobanco://usuario:senha@host:porta/banco_de_dados?sslmode=disable
	dsn := "postgres://admin:secret@localhost:5433/treinamento_go?sslmode=disable"
	// sql.Open valida os parametros  q foi repassado e prepara oobj *sql.DB. Ele so prepara, mais não conect o bd
	db, err := sql.Open("pgx", dsn) // pgx é o nome do driver que foi instalado "github.com/jackc/pgx/v5/stdlib"
	if err != nil {
		return nil, fmt.Errorf("Falha ao abrir a conexão com o bd: %w", err)
	}

	// Verifica se a conexão está viva
	err = db.Ping()
	if err != nil {
		// Fecha a conexão se o ping fala
		db.Close()
		return nil, fmt.Errorf("Falhar ao pingar o bd: %w", err)
	}
	// Caso funcione, a func retorna db e nil para o error (Indicando sucessor)
	return db, nil
}