package app

import (
	"github.com/mkruczek/user-store/config"
	"github.com/mkruczek/user-store/datasource/postgresql"
	"github.com/mkruczek/user-store/utils/logger"
)

func Run() {

	cfg := config.GetApplicationConfig()

	log := logger.New(cfg)

	//db management
	err := postgresql.DoMagicWithDB(cfg, log)
	if err != nil {
		log.Errorf("couldn't managed the DB : %s", err.Error())
	}

	//create nad start webServer
	s := newServer(log)
	mapUrls(s, cfg)
	err = s.start(cfg)
	if err != nil {
		log.Errorf("couldn't start server : %s", err.Error())
	}
}
