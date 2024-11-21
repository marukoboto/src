package main

import "fmt"

//論理演算子

func main(){
	//両方がtrue出ない場合
	fmt.Println(true && false == true) //false
	fmt.Println(true && true == true)	//true
	fmt.Println(true && false == false)	//true
	fmt.Println(false && true  == true)  //false

	//どちらかがtrueの場合true
	fmt.Println(true || false == true) //true
	fmt.Println(false || false == true) //false

	//論理値の反転
	fmt.Println(!true) //false
	fmt.Println(!false) //true

}