package repository

import (
	"encoding/json"
	"os"
	"time"

	cfg "github.com/drenk83/task_tracker_cli/config"
)

type Task struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func OpenFile() (*os.File, error) {
	return os.OpenFile(cfg.StorageFile, os.O_CREATE|os.O_RDWR, 0o644)
}

func WriteTasks(tasks []Task) error {
	tmpName := cfg.StorageFile + ".tmp"
	file, err := os.Create(tmpName)
	if err != nil {
		return err
	}
	defer os.Remove(tmpName)

	enc := json.NewEncoder(file)
	enc.SetIndent("", " ")
	defer file.Close()

	if err := enc.Encode(tasks); err != nil {
		return err
	}

	return os.Rename(tmpName, cfg.StorageFile)
}

func LoadTasks() ([]Task, error) {
	data, err := os.ReadFile(cfg.StorageFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	var taskArray []Task
	err = json.Unmarshal(data, &taskArray)
	if err != nil {
		return nil, err
	}

	return taskArray, nil
}
