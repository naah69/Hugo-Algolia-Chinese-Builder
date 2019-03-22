package utils

import "fmt"

type ThreadPool struct {
	Queue  chan func() error
	Number int
	Total  int

	result         chan error
	finishCallback func()
}

// 初始化
func (self *ThreadPool) Init(number int, total int) {
	self.Queue = make(chan func() error, total)
	self.Number = number
	self.Total = total
	self.result = make(chan error, total)
}

// 开始执行
func (self *ThreadPool) Start() {
	// 开启Number个goroutine
	for i := 0; i < self.Number; i++ {
		go func() {
			for {
				task, ok := <-self.Queue
				if !ok {
					break
				}

				err := task()
				self.result <- err
			}
		}()
	}

	// 获得每个work的执行结果
	for j := 0; j < self.Total; j++ {
		res, ok := <-self.result
		if !ok {
			break
		}

		if res != nil {
			fmt.Println(res)
		}
	}

	// 所有任务都执行完成，回调函数
	if self.finishCallback != nil {
		self.finishCallback()
	}
}

// 停止
func (self *ThreadPool) Stop() {
	close(self.Queue)
	close(self.result)
}

// 添加任务
func (self *ThreadPool) AddTask(task func() error) {
	self.Queue <- task
}

// 设置结束回调
func (self *ThreadPool) SetFinishCallback(callback func()) {
	self.finishCallback = callback
}
