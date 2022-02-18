package comment

import "log"

type Comment interface {
	// All Get latest comments
	All(limit int32, filter, cursor string)
	// Get a comment object
	Get(id, relationships string)
	// Update a comment
	Update(id, data string)
	// Delete a comment
	Delete(id string)
	// Vote Add a vote to a comment
	Vote(id, vote string)
	// Objects Get objects related to a comment
	Objects(id, relationship string)
	// ObjectDescriptors Get object descriptors related to a comment
	/*
		This endpoint is the same as /comments/{id}/{relationship}
		except it returns just the related object's IDs (and context attributes, if any) instead of returning all attributes.
	*/
	ObjectDescriptors(id, relationship, limit, cursor string)
}

type comment struct {
	apiKey string
	log    *log.Logger
}

func New(apiKey string, log *log.Logger) Comment {
	return comment{
		apiKey: apiKey,
		log:    log,
	}
}

func (c comment) All(limit int32, filter, cursor string) {
	//TODO implement me
	panic("implement me")
}

func (c comment) Get(id, relationships string) {
	//TODO implement me
	panic("implement me")
}

func (c comment) Update(id, data string) {
	//TODO implement me
	panic("implement me")
}

func (c comment) Delete(id string) {
	//TODO implement me
	panic("implement me")
}

func (c comment) Vote(id, vote string) {
	//TODO implement me
	panic("implement me")
}

func (c comment) Objects(id, relationship string) {
	//TODO implement me
	panic("implement me")
}

func (c comment) ObjectDescriptors(id, relationship, limit, cursor string) {
	//TODO implement me
	panic("implement me")
}
