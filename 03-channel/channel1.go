package main

import "fmt"

func main() {
	// var channel1 chan int // 宣告 channel 名稱 channel 裡面可以放的種類

	// channel1 = make(chan int, 1) // 設置 channel1 的容量
	// fmt.Println(len(channel1)) // 0
	// fmt.Println(cap(channel1)) //印出 channel1 的容量情況
	// channel1 <- 10
	// fmt.Println(len(channel1)) //查看 channel1 裡面放的情況
	// fmt.Println(<-channel1)    // 印出取出放進 channel1 的數字
	// fmt.Println(len(channel1))

	// channel2 := make(chan int, 3) // 直接宣告 channel 並且給予空間
	// channel2 <- 1
	// channel2 <- 2
	// channel2 <- 3
	// b := <-channel2
	// channel2 <- 4

	// fmt.Println(b) // channel 是先進先出，channel 如果容量滿會被 block

	channel3 := make(chan string) //沒有寫容量代表 0
	// channel3 <- "good"                 // 往沒有容量的 channel 丟資料會被 block
	go func() {
		channel3 <- "hello"
	}()
	fmt.Println(<-channel3) // 但是如果有同時有兩個協程在 channel 相互丟接資料（一來一往），則可以成功，印出 hello
}

/*
重點：
1. var ch chan int 只是宣告，值是 nil，要用 make 才能真的用；對 nil channel 送或收都會永遠卡住。
2. make(chan T, n)：n 是容量（cap），len 是目前放了幾個。沒寫 n 就是 0，也就是 unbuffered channel。
3. channel 是先進先出：channel2 放 1、2、3 後取出的是 1，空出位置才能再放 4。
4. 有容量的 channel：滿了再送會卡住，空了再收也會卡住。
5. 沒容量的 channel：送的一方要等到有人來收才會往下走，所以送和收要在不同 goroutine。
6. 如果卡住的是 main，而且沒有其他 goroutine 能解開，
   程式會直接報 fatal error: all goroutines are asleep - deadlock!
7. <-channel3 會等到拿到值才往下，所以這裡不用 WaitGroup，main 也會等 goroutine 送完。
*/
