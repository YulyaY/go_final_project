package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/domain"
	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

func (h *Handler) DoneTask(w http.ResponseWriter, r *http.Request) {
	var id int
	now := time.Now()
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	if id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: errIdIsEmpty.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	t, err := h.repo.GetTask(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	if t.Repeat == "" {
		err := h.repo.DeleteTask(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
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
		w.Write(result)
		return
	}

	nextDate, err := domain.NextDate(now, t.Date, t.Repeat)
	if err == domain.ErrWrongFormat {
		err := h.repo.DeleteTask(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
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
		w.Write(result)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	t.Date = nextDate
	err = h.repo.UpdateTask(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
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
	w.Write(result)
}
