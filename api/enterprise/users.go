package enterprise

import "log"

type User interface {
	Get(id string)
	Update(id, data string)
	Delete(id string)
	APIUsage(id, startDate, endDate string)
	OverallQuotas(id string)
	Relationship(id, relationship, cursor string, limit int)
	Relationships(id, relationship, cursor string, limit int)
}

type user struct {
	log *log.Logger
}
