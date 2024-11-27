package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var id int
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

	errOfDeleteTask := h.repo.DeleteTask(id)
	if errOfDeleteTask != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: errOfDeleteTask.Error()}.jsonBytes()
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
