package main

import "fmt"

//型
//整数型

func main(){
	//環境依存
	var i int = 100


	// 最大値 最小値
	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// 型指定
	var i2 int64 = 200

	fmt.Println(i + 50)

	// fmt.Println(i + i2)

	fmt.Printf("%T\n", i2)
	//int型（32ビットのint型に変更）
	fmt.Printf("%T\n", int32(i2))
	//int型（環境依存のint型に変更）
	fmt.Println(i + int(i2))
}