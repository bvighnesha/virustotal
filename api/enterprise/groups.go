package enterprise

type Group interface {
	Get(id string)
	Update(id, data string)
	APIUsage(id, startDate, endDate string)
	Administrators(id string)
	UpdateAdministrators(id, data string)
	IsAdministrator(id, userId string)
	RevokeGroupAdminPrivileges(id, userId string)
	Members(id string)
	IsMember(id, userId string)
	Delete(id, userId string)
	Add(id, data string)
	Relationship(id, relationship, limit string, cursor int)
	Relationships(id, relationship, limit string, cursor int)
}

type group struct {
}
