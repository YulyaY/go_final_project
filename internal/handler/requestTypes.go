package handler

type requestTask struct {
	Id      string `json:"id,omitempty"`
	Date    string `json:"date,omitempty"`
	Title   string `json:"title,omitempty"`
	Comment string `json:"comment,omitempty"`
	Repeat  string `json:"repeat,omitempty"`
}

type result struct {
	Id int `json:"id"`
}

const (
	formatDate = "20060102"
)
