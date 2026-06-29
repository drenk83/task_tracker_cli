package service

import (
	"fmt"

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

	return nil
}

func UpdateTask(args []string) error {
	return nil
}

func DeleteTask(args []string) error {
	return nil
}

func ListTasks(args []string) error {
	data, err := r.LoadTasks()
	if err != nil {
		return err
	}
	if len(data) == 0 {
		fmt.Println("No tasks")
	}
	return nil
}
