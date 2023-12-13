package config

import (
	"log"

	"todo_go/utils"

	"gopkg.in/ini.v1" // Import the ini package
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

// LoadConfig は指定された設定ファイル（config.ini）から設定を読み込み、
// ConfigList型のグローバル変数 Config に設定を格納します。
// 設定が正常に読み込まれない場合は、エラーをログに記録してプログラムを終了します。
func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
