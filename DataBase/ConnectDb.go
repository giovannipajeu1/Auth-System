package DataBase

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Conectar() (*sql.DB, error) {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Erro ao carregar o arquivo .env: %w", err)
	}

	// Recupera as variáveis de ambiente para a conexão com o banco de dados
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	charset := os.Getenv("DB_CHARSET")
	parseTime := os.Getenv("DB_PARSE_TIME")
	loc := os.Getenv("DB_LOC")

	// Monta a string de conexão
	stringConnection := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=%s&loc=%s", user, password, dbName, charset, parseTime, loc)

	// Abre a conexão com o banco de dados
	db, erro := sql.Open("mysql", stringConnection)
	if erro != nil {
		return nil, fmt.Errorf("Erro ao abrir a conexão com o banco de dados: %w", erro)
	}

	// Verifica se a conexão está funcionando
	if erro = db.Ping(); erro != nil {
		return nil, fmt.Errorf("Erro ao verificar a conexão com o banco de dados: %w", erro)
	}

	return db, nil

}
