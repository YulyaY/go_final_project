package handler

import (
	"net/http"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/domain"
)

type NextDate struct {
	Now    time.Time `json:"now"`
	Date   string    `json:"date"`
	Repeat string    `json:"repeat"`
}

func (h *Handler) NextDate(w http.ResponseWriter, r *http.Request) {
	var nextDateStruct NextDate
	nowString := r.FormValue("now")
	var err error
	nextDateStruct.Now, err = time.Parse("20060102", nowString)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	nextDateStruct.Date = r.FormValue("date")
	nextDateStruct.Repeat = r.FormValue("repeat")

	nextDate, err := domain.NextDate(nextDateStruct.Now, nextDateStruct.Date, nextDateStruct.Repeat)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(nextDate))
}
