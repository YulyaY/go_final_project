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
	var t model.Task // вызвать т реквест
	now := time.Now()
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	// этот иф в домэйн
	if t.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: errTitleIsEmpty.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	// здесь
	if t.Date == "" {
		t.Date = now.Format(formatDate)
	}

	// парс здесь
	dateParse, err := time.Parse(formatDate, t.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	domain.AddTask(
		model.Task{
			Date: dateParse.Format(domain.FormatDate),
			// ....
		}
	)

	// домэйн
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

	// новая переменная т домэйн
	// и присвоить т домэйн значения атрибутов т реквест

	resultId, err := h.repo.AddTask(t) // обращение к domain

	//это остается
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	res := result{Id: resultId}
	result, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}
