package main

import (
	"fmt"
	booksapiserver "github.com/gdtrivedi/gopractice/booksapi/server"
	"github.com/gdtrivedi/gopractice/forloop"
	"github.com/gdtrivedi/gopractice/hubspotemail"
	"github.com/gdtrivedi/gopractice/stringarray"
	"github.com/gdtrivedi/gopractice/strings"
)

func main() {
	//pkg_stringarray() //test of stringarray package.
	//pkg_forloop() //test of forloop package.
	//pkg_booksapi() //test of booksapi package.
	//pkg_hubspotemail() //test of hubspotemail package.
	pkg_strings() // test of strings package.
}

func pkg_strings() {
	strings.FormatString()
}

func pkg_hubspotemail() {
	client := hubspotemail.NewClient("https://api.hubapi.com", "")

	client.AuthToken()

	/*
	email := "gautam.trivedi@gmail.com"
	err := client.CreateHubspotContact(email)
	if err != nil {
		fmt.Println(err)
	}
	err = client.EnrollContactToWorkflow(email)
	if err != nil {
		fmt.Println(err)
	}
	 */
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
