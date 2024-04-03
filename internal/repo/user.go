package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/BabyJhon/library-managment/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) CreateUser(ctx context.Context, user entity.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, sure_name, phone_number) values ($1, $2, $3) RETURNING id", postgres.UsersTable)

	row := u.db.QueryRow(ctx, query, user.Name, user.SureName, user.PhoneNumber)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	
	return id, nil
}

func (u *UserRepo) DeleteUser(ctx context.Context, id int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", postgres.UsersTable)
	row, err := u.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := row.RowsAffected() //пользователь не найден
	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return err
}

func (u *UserRepo) GetUser(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.UsersTable)

	row := u.db.QueryRow(ctx, query, id)
	if err := row.Scan(&user.Id, &user.Name, &user.SureName, &user.PhoneNumber); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, ErrUserNotFound
		}
		return entity.User{}, err
	}

	return user, nil
}
