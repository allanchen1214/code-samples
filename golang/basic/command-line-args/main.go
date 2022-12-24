package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	str string
	num int
)

func main() {
	// define flags
	flag.StringVar(&str, "str", "default value", "text description")
	flag.IntVar(&num, "num", 5, "text description")

	flag.Parse()

	fmt.Printf("str:%s\n", str)
	fmt.Printf("num:%d\n", num)

	for i, v := range os.Args {
		fmt.Printf("i:%d v:%s\n", i, v)
	}

	fmt.Println("----")

	for i, v := range flag.Args() {
		fmt.Printf("i:%d v:%s", i, v)
	}

	/*
	 * Example output
	   ./test --str=aaaa --num=1000 XX
	   str:aaaa
	   num:1000
	   i:0 v:./test
	   i:1 v:--str=aaaa
	   i:2 v:--num=1000
	   i:3 v:XX
	   ----
	   i:0 v:XX
	*/
}
