package main

import (
	"fmt"
	"os"

	cfg "github.com/drenk83/task_tracker_cli/config"
)

const (
	ADD    = "add"
	UPDATE = "update"
	DELETE = "delete"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		ProgramInstruction()
		os.Exit(0)
	}

	switch args[0] {
	case cfg.Add:
		Add()
	case cfg.Update:
		Update()
	case cfg.Delete:
		Delete()
	case cfg.List:
		List()
	default:
		fmt.Println("Unknow arguments")
	}
}

func ProgramInstruction() {
	fmt.Println("No argumetns")
}

func Add() {
	fmt.Println("Add argument")
}

func Update() {
	fmt.Println("Update argument")
}

func Delete() {
	fmt.Println("Delete argument")
}

func List() {
	fmt.Println("List argument")
}
