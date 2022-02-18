package enterprise

type LiveHunt interface {
	RuleSets()
	Create(data string)
	DeleteAllRuleSets(data, header /*x-confirm-delete*/ string)
	RuleSet(id string)
	Update(id, data string)
	IsRuleSetEditor(id, user_or_group string)
	DeleteRuleSetEditor(id, user_or_group string)
	DeleteRuleSet(id string)
	RuleSetRelationship(id, relationship string)
	UpdateRuleSetRelationship(id, relationship, data string)
	Notifications(limit, filter, cursor string)
	DeleteNotifications(tag string)
	Notification(id string)
	DeleteNotification(id string)
	NotificationsFiles(limit, cursor, filter string)
}

type liveHunt struct {
}
