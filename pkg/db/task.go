package db

import (
	"fmt"
	"time"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// CheckDate — экспортируемая функция для проверки и корректировки даты задачи
func CheckDate(task *Task) error {
	now := time.Now()

	// Если дата пустая, берём сегодняшнюю
	if task.Date == "" {
		task.Date = now.Format("20060102")
		return nil
	}

	// Парсим дату
	t, err := time.Parse("20060102", task.Date)
	if err != nil {
		return fmt.Errorf("дата представлена в формате, отличном от 20060102")
	}

	// Проверяем корректность правила повторения, если оно есть
	if task.Repeat != "" {
		next, err := NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return fmt.Errorf("правило повторения указано в неправильном формате")
		}
	// Если дата в прошлом и есть правило повторения, берём следующую дату
	if afterNow(now, t) {
		task.Date = next
	}
	} else {
	// Если дата в прошлом и нет правила повторения, берём сегодняшнюю дату
		if afterNow(now, t) {
			task.Date = now.Format("20060102")
	}
	}

	return nil
}

// NextDate вычисляет следующую дату выполнения задачи
func NextDate(now time.Time, date, repeat string) (string, error) {
	// Здесь должна быть ваша логика обработки правил повторения
	return date, nil
}

// afterNow проверяет, что дата больше текущей
func afterNow(now, t time.Time) bool {
	return t.Before(now)
}

func AddTask(task *Task) (int64, error) {
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	res, err := db.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
