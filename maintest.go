package main

import (
	"fmt"
	booksapiserver "github.com/gdtrivedi/gopractice/booksapi/server"
	"github.com/gdtrivedi/gopractice/forloop"
	"github.com/gdtrivedi/gopractice/stringarray"
)

func main() {
	pkg_stringarray() //test of stringarray package.
	//pkg_forloop() //test of forloop package.
	//pkg_booksapi() //test of booksapi package.
}

func pkg_stringarray() {
	// START: removeElements test
	inputStrArr2 := stringarray.StrArr{"element1", " element2", " element3 ", "element4 ", "     "}
	selector := ""
	trimElementBeforeMatch := true

	elements2DArr := stringarray.RemoveElements(inputStrArr2, selector, trimElementBeforeMatch)
	fmt.Println("elements2DArr: ", elements2DArr)
	// END: removeElements test
}

func pkg_forloop() {
	forloop.ForloopPractice()
}

func pkg_booksapi() {
	booksapiserver.NewMuxServer()
}
