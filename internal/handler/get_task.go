package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	var id int
	w.Header().Set(contentType, valueJson)
	id, err := strconv.Atoi(r.FormValue(valueId))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	task, err := h.service.GetTask(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	dateParse, err := time.Parse(formatDate, task.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	requestTask := requestTask{
		Id:      task.Id,
		Date:    dateParse.Format(formatDate),
		Title:   task.Title,
		Comment: task.Comment,
		Repeat:  task.Repeat,
	}

	taskToSerialize, err := json.Marshal(requestTask)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(taskToSerialize)
}
