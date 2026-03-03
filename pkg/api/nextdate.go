package api

import (
	"net/http"
	"time"

	"go_final_project/pkg/repeat"
)

const (
	DateFormat = "20060102" // Формат даты, используемый в API
)

// nextDateHandler обрабатывает запросы к /api/nextdate
// URL-параметры:
// - now (опционально): текущая дата в формате "20060102", по умолчанию — текущая дата
// - date (обязательно): исходная дата задачи в формате "20060102"
// - repeat (обязательно): правило повторения (например, "y" или "d 7")
func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из GET‑запроса
	nowStr := r.FormValue("now")
	dateStr := r.FormValue("date")
	repeatStr := r.FormValue("repeat")

	// Если now не указан, берём текущую дату
	var now time.Time
	var err error
	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(DateFormat, nowStr)
		if err != nil {
			http.Error(w, "некорректный формат параметра now", http.StatusBadRequest)
			return
	}
	}

	// Проверяем, что обязательные параметры date и repeat переданы
	if dateStr == "" {
		http.Error(w, "отсутствует обязательный параметр date", http.StatusBadRequest)
		return
	}
	if repeatStr == "" {
		http.Error(w, "отсутствует обязательный параметр repeat", http.StatusBadRequest)
		return
	}

	// Вызываем функцию расчёта следующей даты из пакета repeat
	nextDate, err := repeat.NextDate(now, dateStr, repeatStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Возвращаем результат в формате 20060102
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}

// RegisterNextDateHandler регистрирует обработчик для маршрута /api/nextdate
func RegisterNextDateHandler() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
}
