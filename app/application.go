package app

import "github.com/mkruczek/user-store/datasource/postgresql"

const (
	port = ":3012"
)

func StartApplication() {

	//db management
	postgresql.DoMagicWithDB()

	s := newServer()
	mapUrls(s)
	s.start(port)
}
