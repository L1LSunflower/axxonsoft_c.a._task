package cerrors

import "errors"

var (
	ErrReadBody   = errors.New("can't read body")
	TaskNotExist  = errors.New("task by id does not exist")
	ErrCreateTask = errors.New("failed to create task")
	ErrGetTask    = errors.New("failed to get task")
)
