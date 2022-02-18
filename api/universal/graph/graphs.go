package graph

type Graph interface {
	// Search graphs
	Search(filter, cursor, order, attributes string, limit int32)
	// Create a graph
	Create(rawBody string)
	// Get a graph object
	Get(id string)
	// Update a graph object
	Update(id string, nodes, links []interface{}, graphData interface { /*version string*/
	}, private bool)
	// Delete a graph
	Delete(id string)
	// Comments Get comments on a graph
	Comments(id, limit string, cursor int32)
	// Comment Add a comment to a graph
	Comment(id, data string)
	// GraphViewers Get users and groups that can view a graph
	GraphViewers(id, limit string, cursor int32)
	// AddGraphViewer Grant users and groups permission to see a graph
	AddGraphViewer(id string)
	// IsGraphViewer Check if a user or group can view a graph
	IsGraphViewer(id, userOrGroupId string)
	// DeleteGraphViewer Revoke view permission from a user or group
	DeleteGraphViewer(id string)
	// GraphEditors Get users and groups that can edit a graph
	GraphEditors(id, limit string, cursor int32)
	// AddGraphEditor Grant users and groups permission to edit a graph
	AddGraphEditor(id string)
	// IsGraphEditor Check if a user or group can edit a graph
	IsGraphEditor(id string)
	// DeleteGraphEditor Revoke edit graph permissions from a user or group
	DeleteGraphEditor(id string)
	// Objects Get objects related to a graph
	/*
		Graph objects have number of relationships to other objects.
		As mentioned in the Relationships section, those related objects can be retrieved by sending GET requests to the relationship URL.

		Some relationships are accessible only to users who have access to VirusTotal Intelligence.

		The relationships supported by graph objects are documented in the Graph API object page.
	*/
	Objects(id, relationship, limit string, cursor int32)
	// ObjectDescriptors Get object descriptors related to a graph
	/*
		This endpoint is the same as /graphs/{id}/{relationship}
		except it returns just the related object's IDs (and context attributes, if any) instead of returning all attributes.
	*/
	ObjectDescriptors(id, relationship, limit string, cursor int32)
}
