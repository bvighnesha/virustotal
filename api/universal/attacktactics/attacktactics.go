package attacktactics

type AttackTactic interface {
	// Get an attack tactic object
	Get(id string)
	// Objects Get objects related to an attack tactic
	Objects(id, relationship, cursor, limit int32)
	// ObjectDescriptors Get object descriptors related to an attack tactic
	/*
		This endpoint is the same as /attack_tactics/{id}/{relationship}
		except it returns just the related object's descriptor instead of returning all attributes.
	*/
	ObjectDescriptors(id, relationship, cursor, limit int32)
}
