package enterprise

type RetroHunt interface {
	Jobs(limit, filter, cursor string)
	Create(data string)
	Job(id string)
	Delete(id string)
	AbortJob(id string)
	Relationship(id string)
}

type retroHunt struct {
}
