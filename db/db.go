package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Open подключается к PostgreSQL по DSN. Единая точка входа для продакшена.
func Open(dsn string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

// OpenInMemory возвращает in-memory SQLite для автотестов. Один и тот же подход во всех сервисах.
func OpenInMemory() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
}
