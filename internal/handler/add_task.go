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
	var t model.Task
	now := time.Now()
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
		return
	}

	if t.Title == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, fmt.Sprintf(`{"error": "%s"}`, errTitleIsEmpty.Error()))
		return
	}

	if t.Date == "" {
		t.Date = now.Format(domain.FormatDate)
	}
	dateParse, err := time.Parse(domain.FormatDate, t.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, fmt.Sprintf(`{"error": "%s"}`, "error of parse date"))
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
				fmt.Fprintln(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
				return
			}
			assignDateBuf = nextDate
		}
		t.Date = assignDateBuf
	}

	resultId, err := h.repo.AddTask(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
		return
	}

	res := struct {
		Id int `json:"id"`
	}{Id: resultId}
	result, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
