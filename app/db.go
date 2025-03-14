package app

import (
	"database/sql"
	"fmt"
	"os"

	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	path, err := buildPath()
	if err != nil {
		panic(err)
	}
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		panic(err)
	}

	_, err = createTable(db)

	if err != nil {
		fmt.Println("CREATE TABLE ERROR")
		panic(err)
	}
	return db
}

func buildPath() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dirname = filepath.Join(dirname, ".pomodoro")
	if _, err := os.Stat(dirname); os.IsNotExist(err) {
		err = os.Mkdir(dirname, 0755)
		if err != nil {
			panic(err)
		}
	}
	return fmt.Sprintf("%s/store.db", dirname), nil
}

func createTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS Pomodoro (
        name     TEXT NOT NULL,
        duration     TEXT NOT NULL
    );`

	return db.Exec(sql)
}
