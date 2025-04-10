package file

import (
	"os"
	"testing"
	"time"

	"github.com/BrunoFgR/task-tracker/internal/models"

	"github.com/stretchr/testify/assert"
)

func setupTestFile(t *testing.T) string {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "tasks-*.json")
	assert.NoError(t, err)
	tmpfile.Close()
	return tmpfile.Name()
}

func cleanupTestFile(filename string) {
	os.Remove(filename)
}

func TestNew(t *testing.T) {
	// Test with empty file
	filename := setupTestFile(t)
	defer cleanupTestFile(filename)

	// Write empty array to file
	err := os.WriteFile(filename, []byte("[]"), 0644)
	assert.NoError(t, err)

	storage, err := New(filename)
	assert.NoError(t, err)
	assert.NotNil(t, storage)
	assert.Equal(t, 0, len(storage.content))

	// Test with existing tasks
	tasksJSON := `[{"id":1,"description":"Test Task","status":"TODO","created_at":"2023-01-01T12:00:00Z","updated_at":"2023-01-01T12:00:00Z"}]`
	err = os.WriteFile(filename, []byte(tasksJSON), 0644)
	assert.NoError(t, err)

	storage, err = New(filename)
	assert.NoError(t, err)
	assert.NotNil(t, storage)
	assert.Equal(t, 1, len(storage.content))
	assert.Equal(t, 1, storage.content[0].ID)
}

func TestCreate(t *testing.T) {
	filename := setupTestFile(t)
	defer cleanupTestFile(filename)

	// Write empty array to file
	err := os.WriteFile(filename, []byte("[]"), 0644)
	assert.NoError(t, err)

	storage, err := New(filename)
	assert.NoError(t, err)

	// Test creating a new task
	task := models.Task{
		Description: "New Test Task",
		Status:      models.StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = storage.Create(task)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(storage.content))
	assert.Equal(t, 1, storage.content[0].ID)
	assert.Equal(t, "New Test Task", storage.content[0].Description)

	// Test creating a duplicate task
	err = storage.Create(task)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
}

func TestUpdateByID(t *testing.T) {
	filename := setupTestFile(t)
	defer cleanupTestFile(filename)

	// Create storage with one task
	now := time.Now()
	taskJSON := `[{"id":1,"description":"Test Task","status":"TODO","created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"}]`
	err := os.WriteFile(filename, []byte(taskJSON), 0644)
	assert.NoError(t, err)

	storage, err := New(filename)
	assert.NoError(t, err)

	// Test update with valid ID
	err = storage.UpdateByID(1, "Updated Task")
	assert.NoError(t, err)
	assert.Equal(t, "Updated Task", storage.content[0].Description)

	// Test update with invalid ID
	err = storage.UpdateByID(99, "Should Fail")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")

	// Test update with duplicate description
	storage.content = append(storage.content, models.Task{
		ID:          2,
		Description: "Another Task",
		Status:      models.StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	})

	err = storage.UpdateByID(1, "Another Task")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
}

func TestDeleteByID(t *testing.T) {
	filename := setupTestFile(t)
	defer cleanupTestFile(filename)

	// Create storage with one task
	now := time.Now()
	taskJSON := `[{"id":1,"description":"Test Task","status":"TODO","created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"}]`
	err := os.WriteFile(filename, []byte(taskJSON), 0644)
	assert.NoError(t, err)

	storage, err := New(filename)
	assert.NoError(t, err)

	// Test delete with valid ID
	err = storage.DeleteByID(1)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(storage.content))

	// Test delete with invalid ID
	err = storage.DeleteByID(99)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestUpdateStatusByID(t *testing.T) {
	filename := setupTestFile(t)
	defer cleanupTestFile(filename)

	// Create storage with one task
	now := time.Now()
	taskJSON := `[{"id":1,"description":"Test Task","status":"TODO","created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"}]`
	err := os.WriteFile(filename, []byte(taskJSON), 0644)
	assert.NoError(t, err)

	storage, err := New(filename)
	assert.NoError(t, err)

	// Test update status with valid ID
	err = storage.UpdateStatusByID(1, models.StatusInProgress)
	assert.NoError(t, err)
	assert.Equal(t, models.StatusInProgress, storage.content[0].Status)

	// Test update status with invalid ID
	err = storage.UpdateStatusByID(99, models.StatusDone)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestList(t *testing.T) {
	filename := setupTestFile(t)
	defer cleanupTestFile(filename)

	// Create storage with multiple tasks
	now := time.Now()
	tasksJSON := `[
		{"id":1,"description":"Task 1","status":"TODO","created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"},
		{"id":2,"description":"Task 2","status":"IN_PROGRESS","created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"},
		{"id":3,"description":"Task 3","status":"DONE","created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"}
	]`
	err := os.WriteFile(filename, []byte(tasksJSON), 0644)
	assert.NoError(t, err)

	storage, err := New(filename)
	assert.NoError(t, err)

	// Test listing all tasks
	tasks, err := storage.List()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(tasks))
}

func TestListByStatus(t *testing.T) {
	filename := setupTestFile(t)
	defer cleanupTestFile(filename)

	// Create storage with multiple tasks of different statuses
	now := time.Now()
	tasksJSON := `[
		{"id":1,"description":"Task 1","status":"TODO","created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"},
		{"id":2,"description":"Task 2","status":"IN_PROGRESS","created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"},
		{"id":3,"description":"Task 3","status":"DONE","created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"},
		{"id":4,"description":"Task 4","status":"TODO","created_at":"` + now.Format(time.RFC3339) + `","updated_at":"` + now.Format(time.RFC3339) + `"}
	]`
	err := os.WriteFile(filename, []byte(tasksJSON), 0644)
	assert.NoError(t, err)

	storage, err := New(filename)
	assert.NoError(t, err)

	// Test listing tasks by status
	todoTasks, err := storage.ListByStatus(models.StatusTodo)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(todoTasks))
	for _, task := range todoTasks {
		assert.Equal(t, models.StatusTodo, task.Status)
	}

	inProgressTasks, err := storage.ListByStatus(models.StatusInProgress)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(inProgressTasks))
	assert.Equal(t, models.StatusInProgress, inProgressTasks[0].Status)

	doneTasks, err := storage.ListByStatus(models.StatusDone)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(doneTasks))
	assert.Equal(t, models.StatusDone, doneTasks[0].Status)
}
