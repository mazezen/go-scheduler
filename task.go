package goscheduler

import (
	"context"
	"time"
)

type Task struct {
	ID       int
	Name     string
	ExecTime time.Time
	Job      func(ctx context.Context)
	Ctx      context.Context
	Cancel   context.CancelFunc
}