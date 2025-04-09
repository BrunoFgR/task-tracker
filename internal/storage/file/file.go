package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"slices"

	"github.com/BrunoFgR/task-tracker/internal/models"
)

type Storage struct {
	path      string
	filename  string
	content   []models.Task
	increment func() int
}

// Create instance of Storage
func New(filename string) (*Storage, error) {
	file, err := fileExists(filename)
	if err != nil {
		err := createFile(filename)
		if err != nil {
			return nil, err
		}
	}
	tasks, err := readTasksFromFile(file)
	if err != nil {
		return nil, err
	}
	lastId := 0
	if len(tasks) > 0 {
		lastId = tasks[len(tasks)-1].ID
	}
	return &Storage{
		filename:  filename,
		content:   tasks,
		increment: incrementer(lastId),
	}, nil
}

// Create task
func (s *Storage) Create(taskToCreate models.Task) error {
	for _, task := range s.content {
		if task.Description == taskToCreate.Description {
			return errors.New("It doesn't be able to create a task that already exists.")
		}
	}
	taskWithId := s.addID(taskToCreate)
	s.content = append(s.content, taskWithId)
	err := s.writeFile()
	if err != nil {
		return err
	}
	fmt.Printf("Task added successfully (ID: %d)\n", taskWithId.ID)
	return nil
}

// Update task by ID
func (s *Storage) UpdateByID(id int, descriptionToUpdate string) error {
	idxTask := s.findTaskIndex(id)
	if idxTask == -1 {
		return fmt.Errorf("ID %d not found", id)
	}
	if exist := s.descriptionExists(descriptionToUpdate); exist {
		return fmt.Errorf("description '%s' already exists", descriptionToUpdate)
	}
	s.updateDescription(idxTask, descriptionToUpdate)
	err := s.writeFile()
	if err != nil {
		return err
	}
	fmt.Printf("Task updated successfully (ID: %d)\n", id)
	return nil
}

// Delete task by ID
func (s *Storage) DeleteByID(id int) error {
	idxTask := s.findTaskIndex(id)
	if idxTask == -1 {
		return fmt.Errorf("ID %d not found", id)
	}
	s.content = slices.Delete(s.content, idxTask, idxTask+1)
	err := s.writeFile()
	if err != nil {
		return err
	}
	fmt.Printf("Task deleted successfully (ID: %d)\n", id)
	return nil
}

// Update task status by ID
func (s *Storage) UpdateStatusByID(id int, status models.Status) error {
	idxTask := s.findTaskIndex(id)
	if idxTask == -1 {
		return fmt.Errorf("ID %d not found", id)
	}
	s.updateStatus(idxTask, status)
	err := s.writeFile()
	if err != nil {
		return err
	}
	fmt.Printf("Task status updated successfully (ID: %d)\n", id)
	return nil
}

func (s *Storage) addID(task models.Task) models.Task {
	task.ID = s.increment()
	return task
}

func (s *Storage) findTaskIndex(ID int) int {
	for idx, task := range s.content {
		if task.ID == ID {
			return idx
		}
	}
	return -1
}

func (s *Storage) updateStatus(idx int, status models.Status) {
	s.content[idx].Status = status
	s.content[idx].UpdatedAt = time.Now()
}

func (s *Storage) updateDescription(idx int, description string) {
	s.content[idx].Description = description
	s.content[idx].UpdatedAt = time.Now()
}

func (s *Storage) descriptionExists(description string) bool {
	for _, task := range s.content {
		if task.Description == description {
			return true
		}
	}
	return false
}

func createFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	content := []byte("[]")
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func readTasksFromFile(file *os.File) ([]models.Task, error) {
	defer file.Close()

	var tasks []models.Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return tasks, nil
}

func incrementer(lastId int) func() int {
	count := lastId
	increment := func() int {
		count += 1
		return count
	}
	return increment
}

func fileExists(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	return file, nil
}

func (s *Storage) writeFile() error {
	jsonData, err := json.Marshal(s.content)
	if err != nil {
		return fmt.Errorf("failed to encoded JSON file: %w", err)
	}
	err = os.WriteFile(s.filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}
