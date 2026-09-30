package main // 這是一個可以執行的程式，並從 func main()進行執行

import "fmt"

/*func main(){
	fmt.Println("Hello World,你好 世界")
}
*/

/*func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("現在時間，", time.Now())
}*/

/*func main() {
	fmt.Printf("需要有一個數字 %g\n", math.Pi)
}*/

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 14))
}
