package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

// Глобальная переменная для хранения соединения с БД
var db *sql.DB

// schema содержит SQL‑команды для создания таблицы и индекса
const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL,
    comment TEXT,
    repeat VARCHAR(128)
);

CREATE INDEX idx_scheduler_date ON scheduler (date);
`

// Init инициализирует базу данных: открывает соединение и создаёт таблицу с индексом, если файл БД отсутствует
func Init(dbFile string) error {
	// Проверяем существование файла БД
	_, err := os.Stat(dbFile)
	var install bool
	if err != nil {
		install = true
	}

	// Открываем соединение с БД
	var errOpen error
	db, errOpen = sql.Open("sqlite", dbFile)
	if errOpen != nil {
		return errOpen
	}

	// Если файл БД не существовал, создаём таблицу и индекс
	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil
}