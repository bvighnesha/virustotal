package url

type URLs interface {
	/*
		This endpoint returns an Analysis object descriptor (an identifier for the specific analysis)
		which can be used in the GET /analyses/{id} endpoint to get information about analysis status.

		For analysing an URLs that has been previously submitted to VirusTotal you can use POST /urls/{id}/analyse instead.
	*/
	Analyse(url string)
	// Report Get a URLs analysis report
	Report(id string)
	// ReAnalyse Request a URLs rescan (re-analyze)
	/*
		Returns a Analysis object descriptor which can be used in the GET/analyses/{id} API endpoint to get further information about the analysis status.
	*/
	ReAnalyse(id string)
	// Comments Get comments on a URLs
	Comments(id, cursor string, limit int)
	// Comment Add a comment on a URLs
	/*
		With this endpoint you can post a comment for a given URLs. The body for the POST request must be the JSON representation of a comment object.
		Notice however that you don't need to provide an ID for the object, as they are automatically generated for new comments.

		Any word starting with # in your comment's text will be considered a tag, and added to the comment's tag attribute.
	*/
	Comment(id, data string)
	// Votes Get votes on a URLs
	Votes(id, cursor, limit string)
	// Vote Add a vote on a URLs
	Vote(id string)
	// Objects Get objects related to a URLs
	/*
		URLs objects have number of relationships to other URLs and objects.
		As mentioned in the Relationships section, those related objects can be retrieved by sending GET requests to the relationship URLs.

		Some relationships are accessible only to users who have access to VirusTotal Enterprise package.

		The relationships supported by URLs objects are documented in the URLs API object page.
	*/
	Objects(id, relationship, cursor string, limit int32)
	// ObjectDescriptors Get object descriptors related to a URLs
	ObjectDescriptors(id, relationship, cursor, limit string)
}
