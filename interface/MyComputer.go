package _interface

import "fmt"

type MyComputer struct {
}

func (comp *MyComputer) WriteToFile() {
	fmt.Println("MyComputer WriteToFile")
}

func (comp *MyComputer) WriteToStdout() {
	fmt.Println("MyComputer WriteToFile")
}
