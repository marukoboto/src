package main

import "fmt"

//panic recover

func main(){
	//recoverはpanic によって発生したruntimeエラーから復帰できる機能
	defer func ()  {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	//panic によって実行中のプログラムを強制的に停止
	panic("runtime error") //panicに値は入れる必要がある。
	fmt.Println("Start")
}