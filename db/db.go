package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err) //errを出力してから、プログラムを強制終了
		}
	}
	// 感火曜変数で読み込んだ値をURLに付与
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{}) // DBとのコネクション
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connceted")
	return db
}

// CloseDB は指定された gorm.DB データベース接続を閉じる
// 関数はデータベースの接続プールをクリーンアップ
func CloseDB(db *gorm.DB) {
	// gorm.DB から標準の SQL データベースインターフェースを取得
	sqlDB, _ := db.DB()

	// データベース接続を閉じる
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}
