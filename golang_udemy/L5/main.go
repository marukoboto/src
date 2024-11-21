package main

import "fmt"

//定数

//変数名を大文字にすることで他のパッケージからも呼び出せる
const Pi = 3.14

const (
	URL = "http://xxx.co.jp"
	SiteName = "test"
)

//定義していない部分は定義した部分になる
const (
	A = 1
	B 
	C
	D = "A"
	E
	F
	//111AAAとなる
)

const(
	//iotaは連続する整数の連番を生成する。
	//constのみ使用できる
	c0 = iota
	c1
	c2
)


//変数は型の最大値までしか入力（出力）ができない
// var big int = 9223372036854775807 + 1
// 定数は可能
const big  = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi)

	//一度定義したものは再度定義することができない
	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(big - 1)

	fmt.Println(c0, c1, c2)
}