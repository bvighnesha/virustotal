package monitor

import "os"

type SoftwarePublishers interface {
	Items(filter, cursor string, limit int)
	New(file os.File, path, item string)
	UploadURL()
	Get(id string)
	Delete(id string)
	Configure(id, data string)
	Download(id string)
	DownloadURL(id string)
	Analyses(id, cursor, limit string)
	Owner(id string)
	Comments(id, cursor, limit string)
	Statistics(cursor, limit string)
	Events(cursor, filter string)
}
