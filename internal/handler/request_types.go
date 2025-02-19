package handler

import "time"

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

type GetTasksResp struct {
	Tasks []requestTask `json:"tasks"`
}

type NextDate struct {
	Now    time.Time `json:"now"`
	Date   string    `json:"date"`
	Repeat string    `json:"repeat"`
}

type password struct {
	Password string `json:"password"`
}

type tokenRequest struct {
	Token string `json:"token"`
}

const (
	formatDate  = "20060102"
	valueNow    = "now"
	valueDate   = "date"
	valueRepeat = "repeat"
	contentType = "Content-Type"
	valueJson   = "application/json"
	valueToken  = "token"
	valueId     = "id"
	valueFilter = "%%%s%%"
)
