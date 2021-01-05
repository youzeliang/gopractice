package main

func main() {
	// 新建一个定时任务对象
	// 根据cron表达式进行时间调度，cron可以精确到秒，大部分表达式格式也是从秒开始。
	////crontab := cron.New()  默认从分开始进行时间调度
	//crontab := cron.New() //精确到秒
	////定义定时器调用的任务函数
	//task := func() {
	//	fmt.Println("hello world", time.Now())
	//}
	////定时任务
	//spec := "*/1 * * * * ?" //cron表达式，每五秒一次
	//// 添加定时任务,
	//crontab.AddFunc(spec, task)
	//// 启动定时器
	//crontab.Start()
	//// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	// 根据实际情况进行控制
	// select {} //阻塞主线程停止

	var a []int
	a = append(a, 321)
	aa(a)
	print(a[0])

	var ccc = []int{1, 3, 4}
	aa(ccc)
	print(ccc[0])

}

func aa(xxx []int) {
	xxx[0] = 11111
}
