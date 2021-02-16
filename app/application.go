package app

const (
	port = ":3012"
)

func StartApplication() {

	s := newServer()
	mapUrls(s)
	s.start(port)
}
