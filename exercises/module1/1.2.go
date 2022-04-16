/**
 * @Date 2022/4/16
 * @Name 1.2
 * @VariableName
**/
package module1

import (
	"fmt"
	"time"
)

/**
 * @Date 2022/4/12
 * @Name 1.2
 * @VariableName
**/
/**
基于channel一个简单的单线程生产者消费者模型
队列
长度为10，类型为int
生产
1秒向生产者放入元素，满时可以阻塞
消费
每一秒获取，为空时可以阻塞
*/
func HandleOnePointTwo(queueLen int) {
	queue := make(chan int)
	isDone := make(chan bool)
	defer close(isDone)
	//消费者
	consumer(queue, isDone)
	//生产者
	producer(queue, queueLen)
	fmt.Println("退出！")
}
func producer(q chan int, qLen int) {
	for i := 1; i <= qLen; i++ {
		time.Sleep(time.Second * 1)
		q <- i
	}
	//fmt.Println("发送：", <-q)
}
func consumer(rev <-chan int, status chan bool) {
	go func() {
		//crontabTime := time.NewTimer(time.Second)
		crontabTime := time.NewTicker(time.Second)
		for _ = range crontabTime.C {
			select {
			case <-status:
				fmt.Println("当前队列为空")
				return
			default:
				fmt.Println("接收的Value:", <-rev)
			}
		}
	}()
}
