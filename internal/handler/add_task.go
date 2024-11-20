package handler

import (
	"encoding/json"
	"net/http"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	var t model.Task

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//Проверка на входящие данные
	//Если ошибка, возвращаем 404
	// if err := h.repo.AddTask(); err != nil {
	// 	//Возвращаем 500 ошибку
	// }
	//Возвращаем 200

	result, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(result)
}
