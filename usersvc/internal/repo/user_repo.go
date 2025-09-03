package repo

import (
	"context"

	"github.com/asb19/usersvc/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.UserPublicInfo, error)
	GetUsers(ctx context.Context) ([]model.UserPublicInfo, error)
}

type PostgresUserRepository struct {
	db *pgxpool.Pool
}

func NewPostgresUserRepository(db *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.UserPublicInfo, error) {
	var user model.UserPublicInfo
	err := r.db.QueryRow(ctx,
		`SELECT id, name FROM users WHERE id=$1`, id).
		Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetUsers(ctx context.Context) ([]model.UserPublicInfo, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.UserPublicInfo
	for rows.Next() {
		var u model.UserPublicInfo
		if err := rows.Scan(&u.Id, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
