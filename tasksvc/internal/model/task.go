package model

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	Pending    TaskStatus = "pending"
	InProgress TaskStatus = "inprogress"
	Done       TaskStatus = "done"
)

type Task struct {
	Id          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	AssignedTo  uuid.UUID  `json:"assigned_to"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`

	AssignedUser User `json:"assigned_user,omitempty"`
}

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
