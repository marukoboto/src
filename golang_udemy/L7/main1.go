package main

import "fmt"

//関数

func Plus(x, y int) int {
	return x + y
}

// func Div(x,y int) (int, int) {
// 	//intは切り捨て
// 	q := x/y
// 	r := x % y
// 	return q,r
// }

func Div(x,y int) (q int,r int) {
	//intは切り捨て
	//名前付き返り値を使う場合は(q int,r int)名前付き戻り値として定義されているため、
	//変数を定義（q := x/y)しなくてよい
	q = x/y
	r = x % y
	return 
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func Noreturn(){
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	// i2, i3 := Div(9,4)
	// fmt.Println(i2,i3)

	//変数宣言しているが、片方の変数が不要の場合
	//アンダーバーにして対応する
	i2, _ := Div(9,4)
	//出力する方のみ残す
	fmt.Println(i2)

	i4 := Double(1000)
	fmt.Println(i4)

	//返り値は特に無いため、変数で受け取ることもない
	Noreturn()

}
