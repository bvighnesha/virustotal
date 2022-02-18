package enterprise

import (
	"log"
	"os"
)

type Enterprise interface {
	Users() User
	Groups() Group
	Search() Search
	LiveHunt() LiveHunt
	RetroHunt() RetroHunt
	Zip() Zip
	Feeds() Feeds
}

type enterprise struct {
	apiKey string
	log    *log.Logger
}

func (e *enterprise) Users() User {
	panic("implement me")
}

func (e *enterprise) Groups() Group {
	panic("implement me")
}

func (e *enterprise) Search() Search {
	panic("implement me")
}

func (e *enterprise) LiveHunt() LiveHunt {
	panic("implement me")
}

func (e *enterprise) RetroHunt() RetroHunt {
	panic("implement me")
}

func (e *enterprise) Zip() Zip {
	panic("implement me")
}

func (e *enterprise) Feeds() Feeds {
	panic("implement me")
}

func New(apiKey string) Enterprise {
	return &enterprise{
		apiKey: apiKey,
		log:    log.New(os.Stdout, "virusTotal-enterprise:", log.Ltime|log.Ldate|log.Lshortfile),
	}
}
