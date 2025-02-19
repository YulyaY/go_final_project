package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var id int
	w.Header().Set(contentType, valueJson)
	id, err := strconv.Atoi(r.FormValue(valueId))
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

	errOfDeleteTask := h.service.DeleteTask(id)
	if errOfDeleteTask != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: errOfDeleteTask.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	result, err := json.Marshal(requestTask{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}
