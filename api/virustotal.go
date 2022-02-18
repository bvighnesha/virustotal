package api

import (
	"log"
	"net/http"
	"os"
	"vighnesha.in/virustotal/api/enterprise"
	"vighnesha.in/virustotal/api/monitor"
	"vighnesha.in/virustotal/api/universal"
)

type VirusTotal interface {
	Universal() universal.Universal
	Enterprise() enterprise.Enterprise
	Monitor() monitor.Monitor
}

type virusTotal struct {
	Host   string
	apiKey string
	log    *log.Logger
	client *http.Client
}

func New(apiKey string) VirusTotal {
	return &virusTotal{
		Host:   "https://www.virustotal.com",
		apiKey: apiKey,
		log:    log.New(os.Stdout, "virusTotal:", log.Ldate|log.Ltime|log.Lshortfile),
		client: &http.Client{},
	}
}

func (v *virusTotal) Universal() universal.Universal {
	return universal.New(v.apiKey, v.client)
}

func (v *virusTotal) Enterprise() enterprise.Enterprise {
	return enterprise.New(v.apiKey)
}

func (v *virusTotal) Monitor() monitor.Monitor {
	return monitor.New(v.apiKey)
}
