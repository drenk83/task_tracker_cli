package config

import "errors"

const (
	Add    = "add"
	Update = "update"
	Delete = "delete"
	List   = "list"

	MarkInProgress = "mark-in-progress"
	MarkDone       = "mark-done"
)

const (
	StatusDone       = "done"
	StatusTodo       = "TODO"
	StatusInProgress = "in-progress"
)

const (
	StorageFile = "storage.json"
)

var (
	ErrNoArguments     = errors.New("no arguments")
	ErrInvalidArgument = errors.New("invalid argument")
	ErrInvalidID       = errors.New("invalid id")
	ErrIDDosntExist    = errors.New("no id in list")
)
