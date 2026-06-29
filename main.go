package main

import (
	"fmt"
	"os"

	cfg "github.com/drenk83/task_tracker_cli/config"
	s "github.com/drenk83/task_tracker_cli/service"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		ProgramInstruction()
		os.Exit(0)
	}

	var err error

	switch args[0] {
	case cfg.Add:
		err = s.AddTask(args[1:])
	case cfg.Update:
		err = s.UpdateTask(args[1:])
	case cfg.Delete:
		err = s.DeleteTask(args[1:])
	case cfg.List:
		err = s.ListTasks(args[1:])
	default:
		fmt.Println("Unknow arguments")
	}

	if err != nil {
		fmt.Println(err)
	}
}

func ProgramInstruction() {
	fmt.Println("No args")
}
