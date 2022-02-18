package collections

type Collection interface {
	// New Create a new collection
	/*
		Use this endpoint to create new collections.
		In the request body, send a collection object containing its name, description and the elements it will contain (for URLs you can either use the URL or its ID).
		All IOCs must be described as relationships of a newly created Collection object.
	*/
	New(data string)
	// Get a collection
	Get(id string)
	// Update a collection
	/*
		This endpoint allows updating a collection's attributes (such as name or description) and adding new elements to a collection by using a raw text.
	*/
	Update(id, data string)
	// Delete a collection
	Delete(id string)
	// Comments Get comments on a collection
	Comments(id string)
	// Comment Add a comment to a collection
	/*
		With this endpoint you can post a comment for a given collection.
		The body for the POST request must be the JSON representation of a comment object.
		Notice however that you don't need to provide an ID for the object, as they are automatically generated for new comments.

		Any word starting with # in your comment's text will be considered a tag, and added to the comment's tag attribute.
	*/
	Comment(id, data string)
	// Objects Get objects related to a collection
	Objects(id, relationship, cursor string, limit int32)
	// ObjectDescriptors Get object descriptors related to a collection
	ObjectDescriptors(id, relationship, cursor, limit int32)
	// Add new items to a collection
	/*
		As explained in /collections, for the urls relationship you either use {"type": "url", "url": <a URL string>} or
		{"type": "url", "id": <a sha256 URL identifier>} as object descriptors.
	*/
	Add(id, relationship, data string)
	// DeleteItems Delete items from a collection
	DeleteItems(id, relationship, data string)
}
