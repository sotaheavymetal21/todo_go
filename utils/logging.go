package utils

import (
	"io"
	"log"
	"os"
)

// LoggingSettings はログの設定を行います。指定されたログファイルに出力し、
// 同時に標準出力にもログを表示します。
// logFile: ログを保存するファイルのパス
func LoggingSettings(logFile string) {
	// ログファイルを指定のモードで開く
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		// エラーが発生した場合はエラーをログに記録してプログラムを終了
		log.Fatalln(err)
	}

	// ログの出力先を標準出力と指定されたログファイルに設定
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
