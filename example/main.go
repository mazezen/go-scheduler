package main

import (
	"fmt"
	"time"

	goscheduler "github.com/mazezen/go-scheduler"
)

func main() {
	scheduler := goscheduler.NewScheduler()

	fmt.Println(time.Now())
	execTime1 := time.Now().Add(5 * time.Second)
	scheduler.AddTask(1, "Task 1", execTime1, func() {
		fmt.Println(time.Now(), " Task 1 is running!")
	})

	execTime2 := time.Now().Add(10 * time.Second)
	scheduler.AddTask(2, "Task 2", execTime2, func() {
		fmt.Println(time.Now(), "Task 2 is running with some complex logic!")
	})

	fmt.Println("Scheduler started...")
	scheduler.Run()
}