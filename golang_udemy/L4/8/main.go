package main

import (
	"fmt"
	"strconv"
)

//型変換

func main(){
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// //型を調べる
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fl64 = %T\n", fl64)

	// //フロート型からint型に
	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n" , i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3 = %T\n" , i3)
	// //フロート型からint型に切り替えた場合、小数点以下切り捨ての四捨五入となる
	// fmt.Println(i3)
	// //10

	var s string ="100"
	fmt.Printf("s = %T\n" , s)
	
	//アンダーバーを入れることで2つ目の値を破棄
	//Itoa は "Integer to ASCII" の略
	i, _ := strconv.Atoi(s)
	//strconv.Atoiは文字列型からint型に
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	// //変換できない値を入力した場合（文字列など）
	// i, err := strconv.Atoi("A")
	// //if文で分岐してエラーを出力する(エラーハンドリング)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(i)
	// fmt.Printf("i = %T\n", i)

	var i2 int = 200
	//Atoi は "ASCII to Integer" の略です。
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	var h string = "Hello World"
	//文字をバイト型に
	b := []byte(h)
	fmt.Println(b)

	//バイト型を文字列に
	h2 := string(b)
	fmt.Println(h2)
}