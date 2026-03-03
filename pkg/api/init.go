package api

import "net/http"

func Init() error {
	http.HandleFunc("/api/task", taskHandler)
	// Здесь можно добавить другие обработчики
	return nil
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandler(w, r)
	// Другие методы будут добавлены позже
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
