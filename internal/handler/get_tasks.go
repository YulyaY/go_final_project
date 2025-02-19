package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
	"github.com/YulyaY/go_final_project.git/internal/domain/service"
)

const (
	FormatDateForSearch = "02.01.2006"
	valueSearch         = "search"
)

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, valueJson)
	search := r.FormValue(valueSearch)
	searchDate, errDateParse := time.Parse(FormatDateForSearch, search)
	var taskFilter model.GetTaskFilter
	if errDateParse != nil {
		titleFilter := fmt.Sprintf(valueFilter, search)
		taskFilter.TitleFilter = &titleFilter
	} else {
		dateFilter := service.Format(searchDate)
		taskFilter.DateFilter = &dateFilter
	}

	tasks, err := h.service.GetTasks(taskFilter)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	requestTasks := make([]requestTask, 0, 10)
	for _, task := range tasks {
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
		requestTasks = append(requestTasks, requestTask)
	}

	resultToSerialize := GetTasksResp{Tasks: requestTasks}
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
