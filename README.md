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
Task 1 (Task 1) added, scheduled at 2025-04-23 19:45:20.942743 +0800 CST m=+5.000067168
Task 2 (Task 2) added, scheduled at 2025-04-23 19:45:25.942903 +0800 CST m=+10.000226876
Scheduler started...
Task 1 (Task 1) canceled
Task 1 (Task 1) was canceled, removing
Task 1 removed
Executing task 2 (Task 2) at 2025-04-23 19:45:25.943912 +0800 CST m=+10.001363043
Task 2 is running with complex logic!
Task 2 removed
Main function exiting...
```
