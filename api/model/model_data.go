package model

type Data struct {
	Type          string                   `json:"type,omitempty"`
	ID            string                   `json:"id,omitempty"`
	Links         *Links                   `json:"links,omitempty"`
	Attributes    *Attributes              `json:"attributes,omitempty"`
	Relationships map[string]*Relationship `json:"relationships,omitempty"`
}

type Attributes struct {
}

type Links struct {
	Related string `json:"related,omitempty"`
	Self    string `json:"self,omitempty"`
	Next    string `json:"next,omitempty"`
}

type DatumArr []Datum

type Relationship struct {
	Data  DatumArr `json:"data,omitempty"`
	Meta  *Meta    `json:"meta,omitempty"`
	Links *Links   `json:"links,omitempty"`
}

type Datum struct {
	Error *Error `json:"error,omitempty"`
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
}

type Error struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type Meta struct {
	Cursor string `json:"cursor,omitempty"`
}
