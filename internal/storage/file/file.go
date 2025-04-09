package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/BrunoFgR/task-tracker/internal/models"
)

type Storage struct {
	path      string
	filename  string
	content   []models.Task
	increment func() int
}

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

func (s *Storage) Create(taskToCreate models.Task) error {
	for _, task := range s.content {
		if task.Description == taskToCreate.Description {
			return errors.New("It doesn't be able to create a task that already exists.")
		}
	}
	taskWithId := s.addID(taskToCreate)
	s.content = append(s.content, taskWithId)
	err := s.writeFile(s.content)
	if err != nil {
		return err
	}
	fmt.Printf("Task added successfully (ID: %d)\n", taskWithId.ID)
	return nil
}

func (s *Storage) addID(task models.Task) models.Task {
	task.ID = s.increment()
	return task
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

func (s *Storage) writeFile(content []models.Task) error {
	jsonData, err := json.Marshal(content)
	if err != nil {
		return fmt.Errorf("failed to encoded JSON file: %w", err)
	}
	err = os.WriteFile(s.filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}
