package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, _ := os.Open("foo.txt")
	bs1, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs1))

	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil{
		log.Fatal(err)
	}

	bs2, _ := ioutil.ReadFile("foo.txt")
	fmt.Println(string(bs2))

}
