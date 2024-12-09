package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DoneTask(w http.ResponseWriter, r *http.Request) {
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

	errOfDoneTask := h.service.DoneTask(id)

	// t, err := h.repo.GetTask(id)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
	// 	fmt.Fprintln(w, string(respBytes))
	// 	return
	// }

	// if t.Repeat == "" {
	// 	err := h.repo.DeleteTask(id)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
	// 		fmt.Fprintln(w, string(respBytes))
	// 		return
	// 	}
	// 	result, err := json.Marshal(requestTask{})
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
	// 		fmt.Fprintln(w, string(respBytes))
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write(result)
	// 	return
	// }

	// nextDate, err := service.NextDate(now, t.Date, t.Repeat)
	// if err == service.ErrWrongFormat {
	// 	err := h.repo.DeleteTask(id)
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
	// 		fmt.Fprintln(w, string(respBytes))
	// 		return
	// 	}
	// 	result, err := json.Marshal(requestTask{})
	// 	if err != nil {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
	// 		fmt.Fprintln(w, string(respBytes))
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write(result)
	// 	return
	// } else if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	respBytes := responseErrorWrapper{ErrMsg: err.Error()}.jsonBytes()
	// 	fmt.Fprintln(w, string(respBytes))
	// 	return
	// }
	// t.Date = nextDate
	// err = h.repo.UpdateTask(t)

	if errOfDoneTask != nil {
		w.WriteHeader(http.StatusBadRequest)
		respBytes := responseErrorWrapper{ErrMsg: errOfDoneTask.Error()}.jsonBytes()
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
