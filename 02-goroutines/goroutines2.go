package main

import (
	"fmt"
	"time"
)

func main() {
	num := 1000000
	start := time.Now()
	for i := 0; i < num; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
	end := time.Now()
	fmt.Println(end.Unix()-start.Unix(), "seconds")
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	} else if num == 2 {
		return true
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return false
			}

		}
		return true
	}
}

/*
先不透過goroutines 來執行找質數的程式
總共花費 21秒
*/
