package models

import (
	"database/sql"
)

type User struct {
	ID       int
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Register(email, password string) (int, error) {
	stmt := `INSERT INTO users(email, password) VALUES(?, ?)`
	result, err := m.DB.Exec(stmt, email, password)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
