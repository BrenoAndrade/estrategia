package store

import (
	"context"
	dbsql "database/sql"
	"log"
	"os"
	"time"

	"github.com/go-gorp/gorp"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// SQLStore interface for store
type SQLStore interface {
	GetConn() *gorp.DbMap
	Close()

	Crud() CrudStore
}

// SQLSupplierStore crud
type SQLSupplierStore struct {
	crud CrudStore
}

// SQLSupplier struct
type SQLSupplier struct {
	master           *gorp.DbMap
	stores           SQLSupplierStore
	connectionString string
}

// NewSQLSupplier return new intance
func NewSQLSupplier(connectionString string) *SQLSupplier {
	supplier := &SQLSupplier{
		connectionString: connectionString,
	}

	supplier.initConnection()

	supplier.stores.crud = NewSQLCrudStore(supplier)

	return supplier
}

func setupConnection(connectionString string) *gorp.DbMap {
	db, err := dbsql.Open("mysql", connectionString)

	if err != nil {
		log.Println("[ERRO] Critical error - setupConnection - ", err)
		os.Exit(101)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		log.Println("[ERRO] Critical Error - setupConnection - ", err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8MB4"}}

	return dbmap
}

func (s *SQLSupplier) initConnection() {
	s.master = setupConnection(s.connectionString)
}

// GetConn return connection
func (s *SQLSupplier) GetConn() *gorp.DbMap {
	return s.master
}

// Close db connection
func (s *SQLSupplier) Close() {
	s.master.Db.Close()
}

// Crud return interface
func (s *SQLSupplier) Crud() CrudStore {
	return s.stores.crud
}
