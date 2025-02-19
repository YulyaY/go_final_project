package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/domain/service"
)

func (h *Handler) NextDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(contentType, valueJson)
	var nextDateStruct NextDate
	nowString := r.FormValue(valueNow)
	var err error
	nextDateStruct.Now, err = time.Parse(formatDate, nowString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	nextDateStruct.Date = r.FormValue(valueDate)
	nextDateStruct.Repeat = r.FormValue(valueRepeat)

	nextDate, err := service.NextDate(nextDateStruct.Now, nextDateStruct.Date, nextDateStruct.Repeat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
		fmt.Fprintln(w, string(respBytes))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(nextDate))
}
