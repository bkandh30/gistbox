package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"gistbox.bhavya.net/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
	gists  *models.GistModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "web:Password@20@/gistbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger: logger,
		gists:  &models.GistModel{DB: db},
	}

	logger.Info("Starting server", "addr", *addr)
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
