package main

import "fmt"

type user struct {
	email  string
	name   string
	gender string
	adult  bool
}

func main() {
	stringChan := make(chan string, 3)
	stringChan <- "hello"
	// u := user{
	// 	email:  "123@example.com",
	// 	name:   "kk",
	// 	gender: "male",
	// 	adult:  true,
	// }

	userChan := make(chan user, 1)
	test(userChan)
	fmt.Println(<-userChan) //{123@example.com kk male true}
	fmt.Println(userChan)   //0x10223f4b0070 pointer 記憶體位置
}

func test(userChan chan user) { //傳遞 pointer 不需要再加 *
	u := user{
		email:  "123@example.com",
		name:   "kk",
		gender: "male",
		adult:  true,
	}
	userChan <- u

}
