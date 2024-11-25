package main

import (
	"fmt"
	"time"
)

//go goroutin

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}


func main() {
	//通常プログラムは上から下に実行されるが
	//goをつける（ゴルーチン）をすることで同時に実行することが出来る。
	go sub()
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}