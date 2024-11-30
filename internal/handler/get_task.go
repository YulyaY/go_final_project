package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	var id int
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	task, err := h.repo.GetTask(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	taskToSerialize, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(taskToSerialize)
}
