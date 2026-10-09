package main

type user struct {
	email  string
	name   string
	gender string
	adult  bool
}

// func main() {
// 	stringChan := make(chan string, 3)
// 	stringChan <- "hello"
// 	// u := user{
// 	// 	email:  "123@example.com",
// 	// 	name:   "kk",
// 	// 	gender: "male",
// 	// 	adult:  true,
// 	// }

// 	userChan := make(chan user, 1)
// 	test(userChan)
// 	fmt.Println(<-userChan) //{123@example.com kk male true}
// 	fmt.Println(userChan)   //0x10223f4b0070 pointer 記憶體位置
// }

// func test(userChan chan user) { //傳遞 pointer 不需要再加 *
// 	u := user{
// 		email:  "123@example.com",
// 		name:   "kk",
// 		gender: "male",
// 		adult:  true,
// 	}
// 	userChan <- u

// }

// func main() {
// 	allChan := make(chan interface{}, 3) //interface
// 	allChan <- user{
// 		email:  "xx@gmail.com",
// 		name:   "hello",
// 		gender: "male",
// 		adult:  true,
// 	}
// 	allChan <- "heello world"
// 	allChan <- 1

// 	fmt.Println(<-allChan) // 先進先出
// }

// 單向 channel

func main() {
	sendChan := make(chan<- string, 3) // 只能放的 channel
	getChan := make(<-chan int, 2)     //只能取出
	sendChan <- "good"
	// b := <-sendChan //invalid operation: cannot receive from send-only channel chan<- string sendChan (variable of type chan<- string)

	// getChan <- 2 //./channel2.go:59:2: invalid operation: cannot send to receive-only channel <-chan int getChan (variable of type <-chan int)
}
