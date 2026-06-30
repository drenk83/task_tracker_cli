package service

import (
	"fmt"
	"strconv"
	"time"

	cfg "github.com/drenk83/task_tracker_cli/config"
	r "github.com/drenk83/task_tracker_cli/repository"
)

func AddTask(args []string) error {
	if len(args) == 0 {
		return cfg.ErrNoArguments
	}

	description := args[0]
	if len(args) > 1 {
		fmt.Println("[WARNING] too many arguments:", args[1:])
		fmt.Println("Will be used:", description)
	}

	tasks, err := r.LoadTasks()
	if err != nil {
		return err
	}

	id := 0
	if len(tasks) != 0 {
		for _, task := range tasks {
			taskId, _ := strconv.Atoi(task.ID)
			id = max(id, taskId)
		}
	}

	tasks = append(tasks, *createNewTask(id+1, description))

	if err := r.WriteTasks(tasks); err != nil {
		return err
	}
	return nil
}

func UpdateTask(args []string) error {
	if len(args) == 0 {
		return cfg.ErrNoArguments
	}
	if len(args) != 2 {
		return cfg.ErrInvalidArgument
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return cfg.ErrInvalidID
	}

	tasks, err := r.LoadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		taskID, err := strconv.Atoi(task.ID)
		if err != nil {
			return err
		}

		if taskID == id {
			tasks[i].Description = args[1]
			tasks[i].UpdatedAt = time.Now()

			err := r.WriteTasks(tasks)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return cfg.ErrIDDosntExist
}

func DeleteTask(args []string) error {
	if len(args) == 0 {
		return cfg.ErrNoArguments
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return cfg.ErrInvalidID
	}

	tasks, err := r.LoadTasks()
	if err != nil {
		return err
	}
	for i, task := range tasks {
		taskID, err := strconv.Atoi(task.ID)
		if err != nil {
			return err
		}

		if taskID == id {
			newTasks := make([]r.Task, 0, len(tasks)-1)
			newTasks = append(newTasks, tasks[:i]...)
			if i+1 < len(tasks) {
				newTasks = append(newTasks, tasks[i+1:]...)
			}

			err := r.WriteTasks(newTasks)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return cfg.ErrIDDosntExist
}

func ListTasks(args []string) error {
	data, err := r.LoadTasks()
	if err != nil {
		return err
	}
	if len(data) == 0 {
		fmt.Println("No tasks")
	}

	for _, task := range data {
		fmt.Println(task)
	}
	return nil
}

func createNewTask(id int, description string) *r.Task {
	return &r.Task{
		ID:          strconv.Itoa(id),
		Description: description,
		Status:      "TODO",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
