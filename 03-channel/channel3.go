package main

import "fmt"

func main() {
	closeChan := make(chan int, 3)
	closeChan <- 3
	closeChan <- 2
	close(closeChan)
	// closeChan <- 1 //panic: send on closed channel
	a := <-closeChan //關掉後 chan 不能送，但可以從中取出資料
	fmt.Println(a)
}
