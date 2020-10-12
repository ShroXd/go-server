package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	Cfg *ini.File
)

func LoadMongoDB() string {
	initIni()
	mongo, err := Cfg.GetSection("mongodb")
	if err != nil {
		log.Fatal(err)
	}

	HOST := mongo.Key("HOST").String()
	PORT := mongo.Key("PORT").String()

	return "mongodb://" + HOST + ":" + PORT
}

func initIni() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(err)
	}
}
