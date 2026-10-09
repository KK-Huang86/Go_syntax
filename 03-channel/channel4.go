package main

import "fmt"

// for range 取出 channel 資料，並且配合 close 關閉

func main() {
	num := 100
	intChan := make(chan int, num)
	for i := 0; i < num; i++ {
		intChan <- i
	}

	close(intChan) //後續需要 close 關閉

	for result := range intChan { // 透過 for range 一一將相關的值取出
		fmt.Println(result)

	}

}

/**
如果不加 close 的話， for range 會一直從指定的 channel 取資料
91
92
93
94
95
96
97
98
99
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /Users/kk/Desktop/about_Go/go_syntax/03-channel/channel4.go:14 +0xa8
exit status 2
**/
