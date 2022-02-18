package attacktechnics

type AttackTechnic interface {
	// Get an attack technique object
	Get(id string)
	// Objects Get objects related to an attack technique
	Objects(id, relationship, cursor, limit int32)
	// ObjectDescriptors Get object descriptors related to an attack technique
	/*
		This endpoint is the same as /attack_techniques/{id}/{relationship}
		except it returns just the related object's descriptor instead of returning all attributes.
	*/
	ObjectDescriptors(id, relationship, cursor, limit int32)
}
