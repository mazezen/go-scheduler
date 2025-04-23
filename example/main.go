package main

import (
	"fmt"
	"time"

	goscheduler "github.com/mazezen/go-scheduler"
)

func main() {
	// 创建调度器
	scheduler := goscheduler.NewScheduler()
	fmt.Println("tasks num:", len(scheduler.tasks))

	// 示例任务1：5秒后执行
	execTime1 := time.Now().Add(5 * time.Second)
	scheduler.AddTask(1, "Task 1", execTime1, func() {
		fmt.Println("tasks num:", len(scheduler.tasks))
		fmt.Println("Task 1 is running!")
	})

	// 示例任务2：10秒后执行
	execTime2 := time.Now().Add(10 * time.Second)
	scheduler.AddTask(2, "Task 2", execTime2, func() {
		fmt.Println("tasks num:", len(scheduler.tasks))
		fmt.Println("Task 2 is running with some complex logic!")
	})

	// 启动调度器
	fmt.Println("Scheduler started...")
	scheduler.Run()
	fmt.Println("tasks num:", len(scheduler.tasks))
}