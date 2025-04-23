package goscheduler

import "time"

type Task struct {
	ID       int
	Name     string
	ExecTime time.Time
	Job      func()
}