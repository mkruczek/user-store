package app

import (
	"github.com/mkruczek/user-store/config"
	"github.com/mkruczek/user-store/datasource/postgresql"
	"log"
)

func Run() {

	//load and create Config
	cfgPath, err := config.ParseFlags()
	if err != nil {
		log.Fatalf("couldn't load config path : %s", err.Error())
	}
	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatalf("couldn't create config for application : %s", err.Error())
	}

	//db management
	err = postgresql.DoMagicWithDB(cfg)
	if err != nil {
		log.Fatalf("couldn't managed the DB : %s", err.Error())
	}

	//create nad start webServer
	s := newServer()
	mapUrls(s)
	err = s.start(cfg)
	if err != nil {
		log.Fatalf("couldn't start server : %s", err.Error())
	}
}
