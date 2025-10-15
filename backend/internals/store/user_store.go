package store

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db: db,
	}
}

type UserStore interface{
	CreateUser(*User)(*User, error)
	GetUserById(id int64)(*User, error)
}

func (pg *PostgresUserStore) CreateUser(user *User) (*User, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	query := `INSERT INTO users(firstname, lastname, email, username, password) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	err = tx.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Username, user.Password).Scan(&user.Id)

	if err != nil {
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("Transaction Commit Error : %w", err)
	}
	return user, nil

}

func (pg *PostgresUserStore) GetUserById(id int64) (*User, error) {
	return &User{}, nil
}
