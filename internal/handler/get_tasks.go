package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

type GetTasksResp struct {
	Tasks []model.Task `json:"tasks"`
}

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks, err := h.repo.GetTasks()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
		return
	}

	resultToSerialize := GetTasksResp{Tasks: tasks}
	resp, err := json.Marshal(resultToSerialize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, fmt.Sprintf(`{"error": "%s"}`, err.Error()))
	}
}
