package goscheduler

import (
	"time"
)

type Scheduler struct {
	tasks []*Task
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make([]*Task, 0),
	}
}

func (s *Scheduler) AddTask(id int, name string, execTime time.Time, job func()) {
	task := &Task{
		ID:       id,
		Name:     name,
		ExecTime: execTime,
		Job:      job,
	}
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) Run() {
	for {
		if len(s.tasks) == 0 {
			// fmt.Println("All tasks completed, scheduler stopped.")
			return
		}

		now := time.Now()
		for _, task := range s.tasks {
			// 检查是否到达执行时间
			if now.After(task.ExecTime) || now.Equal(task.ExecTime) {
				// fmt.Printf("Executing task: %s (ID: %d) at %s\n", task.Name, task.ID, now.Format(time.RFC3339))
				task.Job()
				// 从任务列表中移除已执行的任务
				s.removeTask(task.ID)
			}
		}
		// 每秒检查一次
		time.Sleep(1 * time.Second)
	}
}

func (s *Scheduler) removeTask(id int) {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return
		}
	}
}