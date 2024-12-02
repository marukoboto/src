package config

import (
	"log"
	"todo_app/utils"

	"github.com/go-ini/ini"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	//configにconfiginiから読み込んだ値を設定
	Config = ConfigList{
		//Port Sectionに値を設定
		Port: cfg.Section("web").Key("port").MustString("8080"),
		//SQLdriver db,driver
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
