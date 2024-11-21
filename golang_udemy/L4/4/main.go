package main

import "fmt"

//型
//文字型

func main(){
	var s string = "Hello Golang"
	fmt.Println(s)
	fmt.Printf("%T\n" , s)

	var si string = "300"
	fmt.Println(si)
	fmt.Printf("%T\n", si)

	//複数行の文字列を作成
	fmt.Println(`test
	test
		test`)
	
	//クォーテーションをつける方法
	fmt.Println("\"")
	fmt.Printf(`"`)

	//バイトの値で表示
	fmt.Println(s[0])
	//stringを先につける事で文字に
	fmt.Println(string(s[0]))
}