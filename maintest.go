package main

import (
	"fmt"

	"github.com/gdtrivedi/gopractice/gcp/firestore"

	"github.com/gdtrivedi/gopractice/timestamp"

	"github.com/gdtrivedi/gopractice/switchcase"

	"github.com/gdtrivedi/gopractice/responseheaders"

	"github.com/gdtrivedi/gopractice/jsonparsing"

	booksapiserver "github.com/gdtrivedi/gopractice/booksapi/server"
	"github.com/gdtrivedi/gopractice/forloop"
	"github.com/gdtrivedi/gopractice/hubspotemail"
	"github.com/gdtrivedi/gopractice/netpackage"
	"github.com/gdtrivedi/gopractice/stringarray"
	"github.com/gdtrivedi/gopractice/strings"
)

func main() {
	//pkg_stringarray() //test of stringarray package.
	//pkg_forloop() //test of forloop package.
	//pkg_booksapi() //test of booksapi package.
	//pkg_hubspotemail() //test of hubspotemail package.
	//pkg_strings() // test of strings package.
	//pkg_net() //test of netpackage package
	//pkg_jsonparsing()
	//pkg_responseheaders()
	//pkg_switchcase()
	//pkg_showtime()
	pkg_gcp()
}
func pkg_gcp() {
	firestore.PrintDocumentJSON("domain-mapper-service", "application")
}
func pkg_showtime() {
	timestamp.ShowTime()
}

func pkg_switchcase() {
	switchcase.SwitchCaseTest()
}
func pkg_responseheaders() {
	responseheaders.GetResponseHeaders()
}

func pkg_jsonparsing() {
	jsonparsing.UnmarshallTest()
}

func pkg_net() {
	netpackage.TestNetPackage()
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
	//inputStrArr2 := stringarray.StrArr{"element1", " element2", " element3 ", "element4 ", "     "}
	//selector := ""
	//trimElementBeforeMatch := true

	//elements2DArr := stringarray.RemoveElements(inputStrArr2, selector, trimElementBeforeMatch)
	//fmt.Println("elements2DArr: ", elements2DArr)
	// END: removeElements test

	// START: Split Test
	// Given
	inputStrArr := []string{"a", "b", "c", "d", "e"}
	// When
	elements := stringarray.Split(inputStrArr, 2)
	fmt.Println("elements: ", elements)

	inputStrArr = []string{""}
	// When
	elements = stringarray.Split(inputStrArr, 2)
	fmt.Println("elements: ", elements)
	// END: Split Test
}

func pkg_forloop() {
	forloop.ForloopPractice()
}

func pkg_booksapi() {
	booksapiserver.NewMuxServer()
}
