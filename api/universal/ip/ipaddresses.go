package ip

type IPAddress interface {
	// Report Get an IP address report
	Report(ip string)
	// Comments Get comments on an IP address
	Comments(ip, cursor string, limit int32)
	// Comment Add a comment to an IP address
	/*
		With this endpoint you can post a comment for a given IP address.
		The body for the POST request must be the JSON representation of a comment object.
		Notice however that you don't need to provide an ID for the object, as they are automatically generated for new comments.

		Any word starting with # in your comment's text will be considered a tag, and added to the comment's tag attribute.
	*/
	Comment(ip, data string)
	// Votes Get votes on an IP address
	Votes(ip string)
	// Vote Add a vote to an IP address
	Vote(ip string)
	// Objects Get objects related to an IP address
	/*
		IP addresses have number of relationships to other objects.
		As mentioned in the Relationships section, those related objects can be retrieved by sending GET requests to the relationship URL.

		All available relationships are documented in the IP address API object page.
	*/
	Objects(ip, relationship, cursor string, limit int32)
	// ObjectDescriptors Get object descriptors related to an IP address
	/*
		This endpoint is the same as /ip_addresses/{ip}/{relationship}
		except it returns just the related object's IDs (and context attributes, if any) instead of returning all attributes.
	*/
	ObjectDescriptors(ip, relationship, cursor, limit string)
}
