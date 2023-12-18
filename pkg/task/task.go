package task

import "time"

type Task struct {
	ID             int
	ProcessingTime time.Duration
}

func New(id int, processingTime time.Duration) *Task {
	return &Task{
		ID:             id,
		ProcessingTime: processingTime,
	}
}
