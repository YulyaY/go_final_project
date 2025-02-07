package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/domain"
	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

type GetTasksResp struct {
	Tasks []model.Task `json:"tasks"`
}

const (
	FormatDateForSearch = "02.01.2006"
)

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	search := r.FormValue("search")
	searchDate, errDateParse := time.Parse(FormatDateForSearch, search)
	var taskFilter model.GetTaskFilter
	if errDateParse != nil {
		titleFilter := fmt.Sprintf("%%%s%%", search)
		taskFilter.TitleFilter = &titleFilter
	} else {
		dateFilter := domain.Format(searchDate)
		taskFilter.DateFilter = &dateFilter
	}

	tasks, err := h.repo.GetTasks(taskFilter, domain.Limit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	resultToSerialize := GetTasksResp{Tasks: tasks}
	resp, err := json.Marshal(resultToSerialize)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resp)
}
