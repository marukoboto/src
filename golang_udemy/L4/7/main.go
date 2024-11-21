package main

import "fmt"

//interface & nil

func main(){
	//nilとなる
	var x interface{}
	fmt.Println(x)

	//数値を入れる
	x = 1
	fmt.Println(x)

	//どんな値でも入力できる。ただし、計算等は行えない
	x = 3.14
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1,2,3}
	fmt.Println(x)

	// x = 2
	// fmt.Println(x + 3)
}