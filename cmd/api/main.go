package main

import (
	"database/sql"
	"flag"

	"github.com/ddddami/events-go-demo/internal/models"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const version = "1.0.0"

type application struct {
	events *models.EventModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP server port")
	flag.Parse()

	db, err := initDB()
	if err != nil {
		panic(err)
	}

	createTables()

	defer db.Close()

	app := &application{
		events: &models.EventModel{DB: db},
	}

	server := gin.Default()

	registerRoutes(server, app)
	server.Run(*addr)
}

var DB *sql.DB

func initDB() (*sql.DB, error) {
	var err error
	DB, err = sql.Open("sqlite3", "events.db")
	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		DB.Close()
		return nil, err
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	return DB, nil
}

func createTables() {
	stmt := `
CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime TEXT NOT NULL,
	userId INTEGER
)
`
	_, err := DB.Exec(stmt)
	if err != nil {
		panic("Error working with db")
	}
}
