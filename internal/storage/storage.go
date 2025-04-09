package storage

import (
	"github.com/BrunoFgR/task-tracker/internal/models"
)

type TaskStorage interface {
	Create(task models.Task) error
}
