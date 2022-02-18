package monitor

import (
	"log"
	"os"
)

type Monitor interface {
	AntivirusPartners() AntivirusPartners
	SoftwarePublishers() SoftwarePublishers
}

type monitor struct {
	apiKey string
	log    *log.Logger
}

func (m *monitor) AntivirusPartners() AntivirusPartners {
	panic("implement me")
}

func (m *monitor) SoftwarePublishers() SoftwarePublishers {
	panic("implement me")
}

func New(apiKey string) Monitor {
	return &monitor{
		apiKey: apiKey,
		log:    log.New(os.Stdout, "virusTotal-monitor", log.Ltime|log.Ldate|log.Lshortfile),
	}
}
