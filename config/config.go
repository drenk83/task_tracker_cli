package config

import "errors"

const (
	Add    = "add"
	Update = "update"
	Delete = "delete"
	List   = "list"

	MarkInProgress = ""
	MarkDone       = ""

	StatusDone       = ""
	StatusTodo       = ""
	StatusInProgress = ""
)

const (
	StorageFile = "storage.json"
)

var (
	ErrNoArguments     = errors.New("no arguments")
	ErrInvalidArgument = errors.New("invalid argument")
)
