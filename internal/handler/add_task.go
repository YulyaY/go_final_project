package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/domain"
	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

var errTitleIsEmpty error = errors.New("title is empty")

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t model.Task
	now := time.Now()
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	if t.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: errTitleIsEmpty.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	if t.Date == "" {
		t.Date = now.Format(domain.FormatDate)
	}
	dateParse, err := time.Parse(domain.FormatDate, t.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	if !dateParse.After(now) && domain.IsDateNotTheSameDayAsNow(now, dateParse) {
		var assignDateBuf string
		if t.Repeat == "" {
			assignDateBuf = now.Format(domain.FormatDate)
		} else {
			nextDate, err := domain.NextDate(now, t.Date, t.Repeat)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
				fmt.Fprintln(w, string(respBytes))
				return
			}
			assignDateBuf = nextDate
		}
		t.Date = assignDateBuf
	}

	resultId, err := h.repo.AddTask(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	res := struct {
		Id int `json:"id"`
	}{Id: resultId}
	result, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
