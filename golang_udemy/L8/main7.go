package main

import (
	"fmt"
	"os"
)

//defer 関数が終了された時に処理されるもの

func TestDefer() {
	defer fmt.Println("END") //後に出力される
	fmt.Println("START")
}

func RunDefer() {
	//後から登録された順で出力される　今回の場合3,2,1
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}


func main() {
	TestDefer()

	//main関数終了時に実行される
	defer func ()  {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()

	RunDefer()

	file, err :=os.Create("test.text") //1 テキストfile作成
	if err != nil{
		fmt.Println(err) //2　error確認
	}
	defer file.Close() //4 デファー文にしているのでfileが閉じられる

	file.Write([]byte("Hello")) //3　helloと書く
}
