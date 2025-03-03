package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
	"github.com/YulyaY/go_final_project.git/internal/domain/service"
)

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, valueJson)
	var reqT requestTask
	now := time.Now()
	if err := json.NewDecoder(r.Body).Decode(&reqT); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	if reqT.Date == "" {
		reqT.Date = now.Format(formatDate)
	}

	dateParse, err := time.Parse(formatDate, reqT.Date)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	resultId, err := h.service.AddTask(
		model.Task{
			Id:      reqT.Id,
			Date:    dateParse.Format(service.FormatDate),
			Title:   reqT.Title,
			Comment: reqT.Comment,
			Repeat:  reqT.Repeat,
		})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}

	res := result{Id: resultId}
	result, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(result)
}
