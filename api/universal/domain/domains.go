package domain

type Domains interface {
	// Get a domain report
	Get(domain string)
	// Comments Get comments on a domain
	Comments(domain, cursor string, limit int32)
	// Comment Add a comment to a domain
	Comment(domain, data string)
	// Votes Get votes on a domain
	Votes(domain string)
	// Vote Add a vote to a domain
	Vote(domain, data string)
	// Objects Get objects related to a domain
	/*
		Domains objects have number of relationships to other Domains and objects.
		As mentioned in the Relationships section, those related objects can be retrieved by sending GET requests to the relationship URL.

		All available relationships are documented in the domain API object page.
	*/
	Objects(domain, relationship, cursor string, limit int)
	// ObjectDescriptors Get object descriptors related to a domain
	/*
		This endpoint is the same as /domains/{domain}/{relationship} except
		it returns just the related object's IDs (and context attributes, if any) instead of returning all attributes.
	*/
	ObjectDescriptors(domain, relationship, cursor, limit string)
	// DNSResolution Get a DNS resolution object
	/*
		This endpoint retrieves a Resolution object by its ID.
		A resolution object ID is made by appending the IP and the domain it resolves to together.
	*/
	DNSResolution(id string)
}
