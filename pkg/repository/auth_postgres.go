package repository

import (
	"fmt"
	"todo"

	"github.com/jmoiron/sqlx"
)

type AuthPostges struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostges {
	return &AuthPostges{db: db}
}

func (r *AuthPostges) CreateUser(user todo.User) (int, error) {
	var id int

	queryRow := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(queryRow, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
