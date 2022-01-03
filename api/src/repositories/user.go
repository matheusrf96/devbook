package repositories

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repo Users) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(`
		INSERT INTO users (name, username, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	lastInsertId := 0
	err = statement.
		QueryRow(user.Name, user.Username, user.Email, user.Password).
		Scan(&lastInsertId)
	if err != nil {
		return 0, nil
	}

	return uint64(lastInsertId), nil
}
