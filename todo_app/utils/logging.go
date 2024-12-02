package utils

import (
	"io"
	"log"
	"os"
)


func LoggingSettings(logFile string) {
	//os.O_RDWR→ファイルの読み書き　os.O_CREATE→ファイルの作成　os.O_APPEND→ファイルの追記
	//os.OpenFileでconfig.iniから取得したログファイルパスで指定されたファイルを開く
	//ファイルパーミッション（ファイルの権限付与）　0666→書き込みも読み込みもどちらもできるよ
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	//複数の出力先にログを送る　標準出力、ログファイル
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//ログの出力形式を設定する
	//log.Ldate→日付 log.Ltime→時刻 log.Lshortfile→ファイル名
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//ログの出力先を設定する
	log.SetOutput(multiLogFile)
}
