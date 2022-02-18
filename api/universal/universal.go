package universal

import (
	"log"
	"net/http"
	"os"
	"vighnesha.in/virustotal/api/model"
	"vighnesha.in/virustotal/api/universal/comment"
	"vighnesha.in/virustotal/api/universal/domain"
	"vighnesha.in/virustotal/api/universal/file"
	"vighnesha.in/virustotal/api/universal/graph"
	"vighnesha.in/virustotal/api/universal/ip"
	"vighnesha.in/virustotal/api/universal/url"
)

type Universal interface {
	Comments() comment.Comment
	Domains() domain.Domains
	Files() file.Files
	Graphs() graph.Graph
	IPAddresses() ip.IPAddress
	URLs() url.URLs
	// Analysis Get a URLs/file analysis
	Analysis(id string) model.Response
	// Objects Get objects related to an analysis
	/*
		As mentioned in the Relationships section, those related objects can be retrieved by sending GET requests to the relationship URLs.

		Available relationships are described in the analysis object documentation.
	*/
	Objects(id string)
	// ObjectIDs Get object IDs related to an analysis
	/*
		This endpoint is the same as /analyses/{id}/{relationship}
		except it returns just the related object's IDs (and context attributes, if any) instead of returning all attributes.
	*/
	ObjectIDs(id, relationship string)
	// Submission Get a submission object
	Submission(id string) model.Response
	// Operation Get an operation object
	Operation(id string) model.Response
	// Search files, URLs, domains, IPs and tag comments
	Search(query string) model.Response
	// Metadata Get VirusTotal metadata
	/*
		This endpoint returns a dictionary with metadata related to VirusTotal,
		which includes a full list of engines in use, a list of existing privileges, etc.
	*/
	Metadata() model.Response
	// PopularThreatCategories Get a list of popular threat categories
	PopularThreatCategories()
}

type universal struct {
	apiKey string
	log    *log.Logger
	client *http.Client
}

func (c *universal) Comments() comment.Comment {
	return comment.New(c.apiKey, c.log)
}

func (c *universal) Domains() domain.Domains {
	panic("implement me")
}

func (c *universal) Files() file.Files {
	return file.New(c.apiKey, c.log, c.client)
}

func (c *universal) Graphs() graph.Graph {
	panic("implement me")
}

func (c *universal) IPAddresses() ip.IPAddress {
	panic("implement me")
}

func (c *universal) URLs() url.URLs {
	panic("implement me")
}

func (c *universal) Analysis(id string) model.Response {
	panic("implement me")
}

func (c *universal) Objects(id string) {
	//TODO implement me
	panic("implement me")
}

func (c *universal) ObjectIDs(id, relationship string) {
	//TODO implement me
	panic("implement me")
}

func (c *universal) Submission(id string) model.Response {
	panic("implement me")
}

func (c *universal) Operation(id string) model.Response {
	panic("implement me")
}

func (c *universal) Search(query string) model.Response {
	panic("implement me")
}

func (c *universal) Metadata() model.Response {
	panic("implement me")
}

func (c *universal) PopularThreatCategories() {
	//TODO implement me
	panic("implement me")
}

func New(apiKey string, client *http.Client) Universal {
	return &universal{
		apiKey: apiKey,
		log:    log.New(os.Stdout, "virusTotal-universal:", log.Ltime|log.Ldate|log.Lshortfile),
		client: client,
	}
}
