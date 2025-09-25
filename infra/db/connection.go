package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionstring() string {
	return "user=postgres password=admin@123 host=localhost port=5432 dbname=ecommerce sslmode=disable"
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionstring()
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return dbCon, nil
}
