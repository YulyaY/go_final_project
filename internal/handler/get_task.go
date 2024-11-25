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
		fmt.Fprintln(w, fmt.Sprintf(`{"error": "%s"}`, "error of parse id"))
		return
	}
	task, err := h.repo.GetTask(id)
	taskToSerialize, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(taskToSerialize))
}
