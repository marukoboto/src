package main

import "fmt"

//関数
//無名関数

func main() {
	//funcの横に名前が定義されていないため、無名関数
	f := func(x, y int) int {
		return x + y
	}

	i := f(1, 2)
	fmt.Println(i)

	//無名関数
	//引数を最後に持ってきて定義した値を変数で受け取る
	i2 := func(x, y int) int {
		return x + y
	}(1, 2)

	fmt.Println(i2)
}
