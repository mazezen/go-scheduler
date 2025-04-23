# go-secheduler

## 安装

```shell
go get github.com/mazezen/go-scheduler
```

## 使用

```go
func main() {
	// 创建调度器
	scheduler := NewScheduler()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 示例任务1：5秒后执行
	scheduler.AddTask(1, "Task 1", time.Now().Add(5*time.Second), func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Printf("Task 1 canceled: %v\n", ctx.Err())
			return
		default:
			fmt.Println("Task 1 is running!")
		}
	})

	// 示例任务2：10秒后执行
	scheduler.AddTask(2, "Task 2", time.Now().Add(10*time.Second), func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Printf("Task 2 canceled: %v\n", ctx.Err())
			return
		default:
			fmt.Println("Task 2 is running with complex logic!")
		}
	})

	// 启动调度器
	fmt.Println("Scheduler started...")
	go scheduler.Run(ctx)

	// 模拟取消任务
	time.Sleep(2 * time.Second)
	scheduler.CancelTask(1)

	// 等待任务完成
	time.Sleep(15 * time.Second)
	cancel() // 停止调度器
	fmt.Println("Main function exiting...")
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
