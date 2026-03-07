package api

import (
	"net/http"

	"go_final_project/pkg/db"
)

// TasksResp — структура для ответа с списком задач
type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

// tasksHandler обрабатывает GET‑запрос /api/tasks
// Возвращает список ближайших задач в формате JSON
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса — GET
	if r.Method != http.MethodGet {
		writeJSON(w, map[string]string{"error": "метод не поддерживается"})
		return
	}

	// Получаем список задач (ограничение — 50 записей)
	tasks, err := db.Tasks(50)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	// Возвращаем ответ в формате JSON
	writeJSON(w, TasksResp{
		Tasks: tasks,
	})
}
