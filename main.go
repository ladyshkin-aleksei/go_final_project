package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"go_final_project/pkg/db"
	"go_final_project/pkg/api"
)

// Port определяет порт для сервера.
// Сначала проверяется переменная окружения TODO_PORT,
// затем — значение из tests/settings.go,
// по умолчанию — 7540.
var Port = 7540

func init() {

	// Проверяем переменную окружения TODO_PORT
	if portStr := os.Getenv("TODO_PORT"); portStr != "" {
		if port, err := strconv.Atoi(portStr); err == nil {
			Port = port
		}
	}
}


func main() {
	// Инициализация базы данных
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("database initialization error: %v", err)
	}

	// Инициализация API
	api.Init()

	webDir := "./web"
	fileServer := http.FileServer(http.Dir(webDir))
	http.Handle("/", fileServer)

	log.Printf("running the web server on the port %d...\n", Port)
	log.Printf("open it in a browser http://localhost:%d/\n", Port)

	err = http.ListenAndServe(fmt.Sprintf(":%d", Port), nil)
	if err != nil {
		log.Fatalf("server startup error: %v", err)
	}
}