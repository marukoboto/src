package main

import "fmt"

//型
//byte{unit8}型

func main(){
	//バイトで出力
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	//バイトから文字列に
	fmt.Println(string(byteA))

	//文字列をバイト型に
	c := []byte("HI")
	fmt.Println(c)

	// バイト型を文字列に
	fmt.Println(string(c))

	//実行はgo run main.go
}