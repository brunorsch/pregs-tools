package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Compra representa uma compra no sistema
type Compra struct {
	ID          int
	Descricao   string
	Valor       float64
	Data        time.Time
	Categoria   string
	Observacoes string
}

var db *sql.DB

// InitDB inicializa a conexão com o banco de dados SQLite
func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./pregs.db")
	if err != nil {
		return fmt.Errorf("erro ao abrir banco de dados: %v", err)
	}

	// Testa a conexão
	if err = db.Ping(); err != nil {
		return fmt.Errorf("erro ao conectar com banco de dados: %v", err)
	}

	// Cria a tabela se não existir
	if err = createTables(); err != nil {
		return fmt.Errorf("erro ao criar tabelas: %v", err)
	}

	log.Println("Banco de dados inicializado com sucesso")
	return nil
}

// createTables cria as tabelas necessárias
func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS compras (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		descricao TEXT NOT NULL,
		valor REAL NOT NULL,
		data DATETIME NOT NULL,
		categoria TEXT,
		observacoes TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(query)
	return err
}

// CloseDB fecha a conexão com o banco de dados
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

// InserirCompra adiciona uma nova compra ao banco
func InserirCompra(compra Compra) error {
	query := `
	INSERT INTO compras (descricao, valor, data, categoria, observacoes)
	VALUES (?, ?, ?, ?, ?)`

	_, err := db.Exec(query, compra.Descricao, compra.Valor, compra.Data, compra.Categoria, compra.Observacoes)
	return err
}

// ListarCompras retorna todas as compras do banco
func ListarCompras() ([]Compra, error) {
	query := `SELECT id, descricao, valor, data, categoria, observacoes FROM compras ORDER BY data DESC`
	
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var compras []Compra
	for rows.Next() {
		var compra Compra
		err := rows.Scan(&compra.ID, &compra.Descricao, &compra.Valor, &compra.Data, &compra.Categoria, &compra.Observacoes)
		if err != nil {
			return nil, err
		}
		compras = append(compras, compra)
	}

	return compras, nil
}

// BuscarCompraPorID busca uma compra específica pelo ID
func BuscarCompraPorID(id int) (*Compra, error) {
	query := `SELECT id, descricao, valor, data, categoria, observacoes FROM compras WHERE id = ?`
	
	var compra Compra
	err := db.QueryRow(query, id).Scan(&compra.ID, &compra.Descricao, &compra.Valor, &compra.Data, &compra.Categoria, &compra.Observacoes)
	if err != nil {
		return nil, err
	}
	
	return &compra, nil
}

// AtualizarCompra atualiza uma compra existente
func AtualizarCompra(compra Compra) error {
	query := `
	UPDATE compras 
	SET descricao = ?, valor = ?, data = ?, categoria = ?, observacoes = ?
	WHERE id = ?`

	_, err := db.Exec(query, compra.Descricao, compra.Valor, compra.Data, compra.Categoria, compra.Observacoes, compra.ID)
	return err
}

// DeletarCompra remove uma compra do banco
func DeletarCompra(id int) error {
	query := `DELETE FROM compras WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
} 