//Package repository contains all
package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"sync"
)

var (
	once          sync.Once
	sqlConnection *sql.DB
)

func NewSQLConnection() *sql.DB {
	return sqlConnection
}

func LoadSQLConnection() error {
	var err error

	once.Do(func() {
		err = loadSQLConnection()
	})

	return err
}

func loadSQLConnection() error {
	sqlConnection = BD()
	return nil
}

type connection struct {
	db  *sql.DB
	err error
}

func (c connection) Connection() interface{} {
	return c.db
}

func (c connection) Close() error {
	if c.err != nil {
		return c.err
	}

	return c.db.Close()
}

func (c connection) Error() error {
	if c.err != nil {
		return c.err
	}

	c.err = c.db.Ping()
	return c.err
}

func BD() *sql.DB {

	c := struct {
		host     string
		user     string
		password string
		database string
		port     string
		driver   string
	}{
		host:     os.Getenv("DB_HOST"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		database: os.Getenv("DB_NAME"),
		port:     os.Getenv("DB_PORT"),
		driver:   os.Getenv("DB_DRIVER"),
	}

	log.Println("init connection data base...")

	bd, err := sql.Open(c.driver,
		fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			c.user,
			c.password,
			c.host,
			c.port,
			c.database,
			"disable"),
	)

	if err != nil {
		panic(err.Error())
	}
	log.Println("connect to data base!")

	return bd
}
