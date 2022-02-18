package monitor

type AntivirusPartners interface {
	Hashes(filter, cursor, limit string)
	Analyses(sha256 string)
	Items(sha256 string)
	CommentWithHash(sha256, data string)
	Comments(id string)
	Comment(id, data string)
	DeleteComment(id string)
	Download(sha256 string)
	DownloadURL(sha256 string)
	DownloadDailyDetections(engine_name string)
	DailyDetectionsDownloadURL(engine_name string)
	Statistics()
}
