package model

type Task struct {
	Id      string
	Date    string
	Title   string
	Comment string
	Repeat  string
}

type GetTaskFilter struct {
	TitleFilter *string
	DateFilter  *string

}
