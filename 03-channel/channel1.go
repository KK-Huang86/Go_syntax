package main

import (
	"fmt"
)

func main() {
	var channel1 chan int // 宣告 channel 名稱 channel 裡面可以放的種類

	channel1 = make(chan int, 1) // 設置 channel1 的容量
	fmt.Println(len(channel1))
	fmt.Println(cap(channel1)) //印出 channel1 的容量情況
	channel1 <- 10
	fmt.Println(len(channel1)) //查看 channel1 裡面放的情況
	fmt.Println(<-channel1)    // 印出取出放進 channel1 的數字
	fmt.Println(len(channel1))

}
