package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/domain"
	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

var errIdIsEmpty error = errors.New("id is empty")
var errIdHasWrongFormat error = errors.New("id has wrong format")

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t model.Task
	now := time.Now()

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	if t.Id == "" {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: errIdIsEmpty.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	if _, err := strconv.Atoi(t.Id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: errIdHasWrongFormat.Error()}.jsonBytes()
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
		t.Date = now.Format(formatDate)
	}
	dateParse, err := time.Parse(formatDate, t.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	if !dateParse.After(now) && domain.IsDateNotTheSameDayAsNow(now, dateParse) {
		var assignDateBuf string
		if t.Repeat == "" {
			assignDateBuf = now.Format(formatDate)
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

	errOfUpdate := h.repo.UpdateTask(t)
	if errOfUpdate != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: errOfUpdate.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	result, err := json.Marshal(model.Task{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}
