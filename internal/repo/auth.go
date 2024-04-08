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

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{db: db}
}

func (a *AuthRepo) CreateAdmin(ctx context.Context, input entity.Admin) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, sure_name,user_name, password_hash) values ($1, $2, $3, $4) RETURNING ID", postgres.AdminsTable)

	row := a.db.QueryRow(ctx, query, input.Name, input.SureName, input.UserName, input.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (a *AuthRepo) GetAdmin(ctx context.Context, userName, password string) (entity.Admin, error) {
	var admin entity.Admin

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_name = $1 AND password_hash = $2", postgres.AdminsTable)

	row := a.db.QueryRow(ctx, query, userName, password)
	if err := row.Scan(&admin.Id, &admin.Name, &admin.SureName, &admin.UserName, &admin.Password); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Admin{}, ErrAdminnotFound
		}
		return entity.Admin{}, err
	}

	return admin, nil
}
