package server

import (
	"go_final_project/pkg/api"
	"log"
	"net/http"
)

// Run запускает HTTP‑сервер на порту 7540 и возвращает ошибку
func Run() error {
	// Регистрируем API‑обработчики
	api.Init()

	log.Println("Сервер запущен на :7540")
	return http.ListenAndServe(":7540", nil)
}
