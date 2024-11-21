package main

import "fmt"

//関数
//関数を返す関数

//無名関数funcに値が返される
func ReturnFunc() func(){
	return func ()  {
		fmt.Println("I'm a function")
	}

}

func main(){
	f:= ReturnFunc()
	f()
}