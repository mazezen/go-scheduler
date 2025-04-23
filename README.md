# go-secheduler

## 安装

```shell
go get github.com/mazezen/go-scheduler
```

## 使用

```go
func main() {
	scheduler := NewScheduler()

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

```

## 结果

```markdown
tasks num: 0
Scheduler started...
Executing task: Task 1 (ID: 1) at 2025-04-23T13:53:01+08:00
tasks num: 2
Task 1 is running!
Executing task: Task 2 (ID: 2) at 2025-04-23T13:53:06+08:00
tasks num: 1
Task 2 is running with some complex logic!
All tasks completed, scheduler stopped.
tasks num: 0
```
