package main

import "fmt"

//関数
//関数を引数に取る関数

//mainで実行されたものを受け取り、実行する。
func CallFunction(f func()){
	f()
}

func main() {
	CallFunction(func ()  {
		fmt.Println("I'm a function")
	})

}