package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Zeta-Manu/manu-lesson/config"
)

type Database struct {
	Conn *sql.DB
}

// NewDatabase creates a new MySQL database connection.
func NewDatabase(dataSourceName string) (*Database, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Check if the database connection is alive
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{Conn: db}, nil
}

// InitializeDatabase initializes and returns a new database connection.
func InitializeDatabase(dbConfig config.DatabaseConfig) (*Database, error) {
	dbDataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
	return NewDatabase(dbDataSourceName)
}

// Close closes the database connection.
func (db *Database) Close() error {
	if db.Conn != nil {
		return db.Conn.Close()
	}
	return nil
}

type DBAdapter interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// Executes a SQL query and returns a result
func (db *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Execute a SQL statement
func (db *Database) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Conn.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
