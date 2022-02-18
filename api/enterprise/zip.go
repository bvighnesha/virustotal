package enterprise

type Zip interface {
	Zip(data Data)
	Get(id string)
	DownloadURL(id string)
	Download(id string)
}

type zip struct {
}

type Data struct {
	Password *string  `json:"password,omitempty"`
	Hashes   []string `json:"hashes,omitempty"`
}
