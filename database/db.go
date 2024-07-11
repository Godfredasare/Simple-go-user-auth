package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := os.Getenv("POSGRES_CONNECTION_STRING")
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB is connected Successfully")

	createTables()
}

func createTables() {
	createUser := `
	CREATE TABLE if NOT EXISTS Users(
    id SERIAL PRIMARY KEY NOT NULL,
	username VARCHAR(120) NOT NULL,
	email VARCHAR NOT NULL UNIQUE,
	Password VARCHAR NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
)
	`

	_, err := DB.Exec(createUser)
	if err != nil {
		log.Fatal("Counld not user table", err)
	}

	createCategories := `
	   CREATE TABLE IF NOT EXISTS category(
	     id SERIAL PRIMARY KEY,
		 name VARCHAR(100),
		 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		 updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP 
	   )
	`
	_, err = DB.Exec(createCategories)
	if err != nil {
		log.Fatal("Counld not create category table", err)
	}

	createProducts := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(120) NOT NULL,
		description TEXT,
		price NUMERIC(10,2) NOT NULL,
		currency VARCHAR(6) NOT NULL DEFAULT 'USD',
		quantity INTEGER NOT NULL DEFAULT 0,
		active BOOLEAN NOT NULL DEFAULT true,
		category_id INTEGER NOT NULL REFERENCES category(id),
		user_id INTEGER NOT NULL REFERENCES users(id),
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = DB.Exec(createProducts)
	if err != nil {
		log.Println(err)
		log.Fatal("Counld not create products table")
	}

}
