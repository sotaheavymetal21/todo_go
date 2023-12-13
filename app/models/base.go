package models

import (
	"database/sql"
	"fmt"
	"log"
	"todo_go/config"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

const (
	tableNameUser = "users"
)

func init() {
	// データベースの初期化
	Db, err := sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	// ユーザーテーブルの作成
	cmdU := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s("+
		"id INTEGER PRIMARY KEY AUTOINCREMENT,"+
		"uuid STRING NOT NULL UNIQUE,"+
		"name STRING,"+
		"email STRING,"+
		"password STRING,"+
		"created_at DATETIME)", tableNameUser)

	_, err = Db.Exec(cmdU)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}
