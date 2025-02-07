package handler

import (
	"fmt"
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
	w.Header().Set("Content-Type", "application/json")
	var nextDateStruct NextDate
	nowString := r.FormValue("now")
	var err error
	nextDateStruct.Now, err = time.Parse(formatDate, nowString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	nextDateStruct.Date = r.FormValue("date")
	nextDateStruct.Repeat = r.FormValue("repeat")

	nextDate, err := domain.NextDate(nextDateStruct.Now, nextDateStruct.Date, nextDateStruct.Repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(nextDate))
}
