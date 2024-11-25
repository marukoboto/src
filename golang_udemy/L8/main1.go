package main

import "fmt"

//if
//条件分岐

func main() {
	a := 1
	if a == 2{
		fmt.Println("two")
	} else if a == 1{
		fmt.Println("one")
	} else {
		fmt.Println("I don't Know")
	}

	if b := 100; b == 100 {
		fmt.Println("one hundred")
	}

	x := 0
	if x := 2; true{
		//if文内なので表示されるのは２
		fmt.Println(x)
	}
	//if文外なので表示されるのは０
	fmt.Println(x)
}