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

func (repo Users) GetUser(Id uint64) (models.User, error) {
	rows, err := repo.db.Query(`
		SELECT id, name, username, email, created_at
		FROM users
		WHERE id = $1
	`, Id)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	var user models.User

	if rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo Users) Update(Id uint64, user models.User) error {
	statement, err := repo.db.Prepare(`
		UPDATE users
		SET name = $1, username = $2, email = $3
		WHERE id = $4
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(user.Name, user.Username, user.Email, Id)
	if err != nil {
		return err
	}

	return nil
}

func (repo Users) Delete(Id uint64) error {
	statement, err := repo.db.Prepare(`
		DELETE FROM users
		WHERE id = $1
	`)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(Id)
	if err != nil {
		return err
	}

	return nil
}

func (repo Users) GetUserByEmail(email string) (models.User, error) {
	row, err := repo.db.Query(`
		SELECT id, password
		FROM users
		WHERE email = $1
	`, email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		err = row.Scan(&user.Id, &user.Password)
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
