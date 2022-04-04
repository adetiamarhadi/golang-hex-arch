package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

// Adapter ...
type Adapter struct {
	db *sql.DB
}

// NewAdapter ...
func NewAdapter(driverName, datasourceName string) (*Adapter, error) {

	db, err := sql.Open(driverName, datasourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

// CloseDbConnection ...
func (a Adapter) CloseDbConnection() {
	err := a.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

// AddToHistory ...
func (a Adapter) AddToHistory(value int32, operation string) error {

	queryString, args, err := sq.Insert("arith_history").Columns("date", "answer", "operation").Values(time.Now(), value, operation).ToSql()
	if err != nil {
		return err
	}

	_, err = a.db.Exec(queryString, args...)
	if err != nil {
		return err
	}

	return nil
}
