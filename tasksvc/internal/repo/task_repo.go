package repo

import (
	"context"
	"time"

	"github.com/asb19/tasksvc/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) (model.Task, error)
	GetAll(ctx context.Context, status string, page, limit int) ([]model.Task, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Task, error)
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type PostgresTaskRepository struct {
	db *pgxpool.Pool
}

func NewPostgresTaskRepository(db *pgxpool.Pool) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) Create(ctx context.Context, task *model.Task) (model.Task, error) {

	var createdTask model.Task

	err := r.db.QueryRow(ctx,
		`INSERT INTO tasks (title, description, status) 
         VALUES ($1, $2, $3) 
         RETURNING id, title, description, status`,
		task.Title, task.Description, task.Status,
	).Scan(&createdTask.Id, &createdTask.Title, &createdTask.Description, &createdTask.Status)

	return createdTask, err
}

func (r *PostgresTaskRepository) GetAll(ctx context.Context, status string, page, limit int) ([]model.Task, error) {
	var rows pgx.Rows
	var err error

	if status != "" {
		rows, err = r.db.Query(ctx,
			`SELECT id, title, description, status FROM tasks WHERE status=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`,
			status, limit, (page-1)*limit,
		)
	} else {
		rows, err = r.db.Query(ctx,
			`SELECT id, title, description, status FROM tasks ORDER BY created_at DESC LIMIT $1 OFFSET $2`,
			limit, (page-1)*limit,
		)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.Id, &t.Title, &t.Description, &t.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *PostgresTaskRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	var task model.Task
	err := r.db.QueryRow(ctx,
		`SELECT id, title, description, status, created_at, updated_at,assigned_to FROM tasks WHERE id=$1`, id).
		Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt, &task.AssignedTo)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *PostgresTaskRepository) Update(ctx context.Context, task *model.Task) error {
	task.UpdatedAt = time.Now()
	_, err := r.db.Exec(ctx,
		`UPDATE tasks SET title=$1, description=$2, status=$3, updated_at=$4 WHERE id=$5`,
		task.Title, task.Description, task.Status, task.UpdatedAt, task.Id,
	)
	return err
}

func (r *PostgresTaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.Exec(ctx, `DELETE FROM tasks WHERE id=$1`, id)
	return err
}
