package context

import "github.com/BrunoFgR/task-tracker/internal/storage"

// AppContext holds application-wide dependencies
type AppContext struct {
	Storage storage.TaskStorage
}

// New creates a new application context
func New(storage storage.TaskStorage) *AppContext {
	return &AppContext{
		Storage: storage,
	}
}
