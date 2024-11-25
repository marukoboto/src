package main

import "fmt"

//init 
//mainよりも後に書いてあっても先に実行されるもの
//複数に分けるメリットはあまりない

func main() {
	fmt.Println("Main")
}

func init(){
	fmt.Println("init")
}

func init(){
	fmt.Println("init2")
}