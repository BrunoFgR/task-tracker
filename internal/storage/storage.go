package storage

import (
	"github.com/BrunoFgR/task-tracker/internal/models"
)

type TaskStorage interface {
	Create(task models.Task) error
	Update(task models.Task) error
	Delete(id string) error
	GetByID(id string) (models.Task, error)
	ListAll() ([]models.Task, error)
	ListByStatus(status models.Status) ([]models.Task, error)
}
