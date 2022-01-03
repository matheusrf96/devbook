package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repo Users) Get(nameOrUser string) ([]models.User, error) {
	nameOrUser = fmt.Sprintf("%%%s%%", nameOrUser)

	rows, err := repo.db.Query(`
		SELECT id, name, username, email, created_at
		FROM users
		WHERE name LIKE $1
			OR username LIKE $2
	`, nameOrUser, nameOrUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
