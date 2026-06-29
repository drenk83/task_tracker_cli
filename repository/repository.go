package repository

import (
	"encoding/json"
	"os"
	"time"

	cfg "github.com/drenk83/task_tracker_cli/config"
)

type Task struct {
	ID          string        `json:"id"`
	Description string        `json:"description"`
	Status      string        `json:"status"`
	CreatedAt   time.Duration `json:"createdAt"`
	UpdatedAt   time.Duration `json:"updatedAt"`
}

func OpenFile() (*os.File, error) {
	return os.OpenFile(cfg.StorageFile, os.O_CREATE|os.O_RDWR, 0o644)
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
