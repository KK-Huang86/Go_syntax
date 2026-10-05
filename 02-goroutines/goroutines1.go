package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
		fmt.Println(s)
	}

}

func main() {
	go say("hello")
	say("world")
}

/*
1. go say("hello")：開一個新的 goroutine 執行 say("hello")
2. main goroutine 繼續執行 say("world")
3.兩邊的執行順序沒有保證，兩個 goroutine 都在競爭執行機會，因此有可能是先 hello 再 world，也有可能相反
4.main 結束時，整個程式直接結束，不會自動等其他 goroutine
*/
