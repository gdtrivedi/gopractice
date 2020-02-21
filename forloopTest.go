package main

import (
	"fmt"

	hello1 "git.gogoair.com/bsg/gocomponents/hello"
	"git.gogoair.com/bsg/gocomponents/math"
	"github.com/gdtrivedi/gomodules/hello"
	"github.com/gdtrivedi/gopractice/forloop"
)

func main() {
	forloop.ForloopPractice()
	fmt.Println(hello.Hello())
	fmt.Println(hello1.Hello())
	fmt.Println(math.Add(1, 2))
}
