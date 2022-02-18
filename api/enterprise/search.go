package enterprise

type Search interface {
	Search()
	Snippets(snippet string)
}

type search struct {
}
