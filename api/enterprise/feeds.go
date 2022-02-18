package enterprise

type Feeds interface {
	Files(time string)
	DownloadFile(token string)
	URLs(time string)
	FileSandboxBehaviour(time string)
	PCAP(token string)
	HTMLReport(token string)
}

type feeds struct {
}
