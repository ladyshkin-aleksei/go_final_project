package api

import (
	"net/http"
	"time"
)

const (
	DateFormat = "20060102" // Формат даты, используемый в API
)

// nextDateHandler обрабатывает запросы к /api/nextdate
func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем параметры из GET-запроса
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

	// Вызываем функцию расчёта следующей даты
	nextDate, err := repeat.NextDate(now, dateStr, repeatStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Возвращаем результат в формате 20060102
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}
