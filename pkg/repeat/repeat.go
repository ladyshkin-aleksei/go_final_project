package task

import (
	"strconv"
	"strings"
	"time"
)

// AddTask добавляет задачу с проверкой валидности данных.
// Возвращает true, если задача валидна и может быть добавлена, иначе false.
func AddTask(now time.Time, date string, title string, repeat string) bool {
	// Парсим дату задачи
	taskDate, err := time.Parse("20060102", date)
	if err != nil {
		return false
	}

	// Нормализуем now и taskDate до начала дня (убираем время)
	nowTrunc := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	taskDateTrunc := time.Date(taskDate.Year(), taskDate.Month(), taskDate.Day(), 0, 0, 0, 0, taskDate.Location())

	// Проверяем, что дата задачи не раньше текущей даты
	if taskDateTrunc.Before(nowTrunc) {
		return false
	}

	// Если правило повторения пустое — валидно
	if repeat == "" {
		return true
	}

	// Разбиваем правило на части
	parts := strings.Split(repeat, " ")

	switch parts[0] {
	case "y":
		// Правило 'y' должно быть единственным словом
		if len(parts) != 1 {
			return false
		}
		return true

	case "d":
		// Правило 'd' должно иметь ровно два слова: 'd' и число
		if len(parts) != 2 {
			return false
		}

		days, err := strconv.Atoi(parts[1])
		if err != nil {
			return false
		}

		if days <= 0 {
			return false
		}
		if days > 400 {
			return false
		}
		return true

	default:
		// Любое другое правило — ошибка
		return false
	}
}
