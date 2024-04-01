package repo

import (
	"fmt"

	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/BabyJhon/library-managment/pkg/postgres"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) CreateUser(user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, sure_name, phone_number) values ($1, $2, $3) RETURNING id", postgres.UsersTable)
	row := u.db.QueryRow(query, user.Name, user.SureName, user.PhoneNumber)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (u *UserRepo) DeleteUser(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", postgres.UsersTable)
	_, err := u.db.Exec(query, id)
	return err
}

func (u *UserRepo) GetUser(id int) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.UsersTable)
	err := u.db.Get(&user, query, id)
	return user, err
}
