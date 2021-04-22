package main

import (
	"fmt"

	"github.com/gdtrivedi/gopractice/fileutil"

	"github.com/gdtrivedi/gopractice/jsontomap"

	"github.com/gdtrivedi/gopractice/uploadfiles"

	"github.com/gdtrivedi/gopractice/certificateparse"

	"github.com/gdtrivedi/gopractice/structtojson"

	"github.com/gdtrivedi/gopractice/maptest"

	"github.com/gdtrivedi/gopractice/array"

	"github.com/gdtrivedi/gopractice/cfclient/gi"

	"github.com/gdtrivedi/gopractice/modeltest"

	"github.com/gdtrivedi/gopractice/strconv"

	"github.com/gdtrivedi/gopractice/split"

	"github.com/gdtrivedi/gopractice/cflogpush/bqwrite"

	"github.com/gdtrivedi/gopractice/base64"

	"github.com/gdtrivedi/gopractice/waitgroup"

	"github.com/gdtrivedi/gopractice/entitlementsvc/client"

	"github.com/gdtrivedi/gopractice/errorf"

	"github.com/gdtrivedi/gopractice/dmsclient"

	"github.com/gdtrivedi/gopractice/querystring"

	"github.com/gdtrivedi/gopractice/jsontostruct"

	"github.com/gdtrivedi/gopractice/uuid"

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
	//pkg_gcp()
	//pkg_uuid()
	//pkg_jsontostruct()
	//pkg_querystring()
	//pkg_dmsclient()
	//pkg_entitlementclient()
	//pkg_errorf()
	//pkg_waitgroup()
	//pkg_base64()
	//pkg_bqwrite()
	//pkg_split()
	//pkg_strconv()
	//pkg_modeltest()
	//pkg_gicftest()
	//pkg_structtojson()
	//pkg_arraytest()
	//pkg_map()
	pkg_certparse()
	//pkg_uploadfiletest()
	//pkg_jsontomap()
	//pkg_fileutil()
}
func pkg_fileutil() {
	fileutil.WriteStringToFile()
}
func pkg_jsontomap() {
	jsontomap.JsontoMapTest()
}
func pkg_uploadfiletest() {
	uploadfiles.UploadFilesTest()
}
func pkg_certparse() {
	certificateparse.ParseCert()
}
func pkg_map() {
	maptest.MapTest()
}
func pkg_arraytest() {
	array.ArrayTest()
}
func pkg_structtojson() {
	//structtojson.TestStructToJson()
	structtojson.TestDNSRecordSpecStructToJson()
}
func pkg_gicftest() {
	gi.GICFTest()
}
func pkg_modeltest() {
	modeltest.ModelTest()
}
func pkg_strconv() {
	strconv.ParseFloatTest()
}
func pkg_split() {
	split.SplitTest()
}
func pkg_bqwrite() {
	bqwrite.BQWriteTest()
}
func pkg_base64() {
	base64.Base64Test()
}
func pkg_waitgroup() {
	waitgroup.WaitgroupSimpleTest()
}
func pkg_entitlementclient() {
	client.EntitlementClientTest()
}
func pkg_errorf() {
	errorf.ErrorfTest()
}
func pkg_dmsclient() {
	dmsclient.DMSClientTest()
}
func pkg_querystring() {
	querystring.QueryStringTest()
}
func pkg_jsontostruct() {
	jsontostruct.InstallDomainsTest()
}
func pkg_uuid() {
	uuid.UUIDTest()
}
func pkg_gcp() {
	firestore.PrintDocumentJSON("domain-mapper-service", "application")
}
func pkg_showtime() {
	//timestamp.ShowTime()
	timestamp.TimeInitTest()
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

	stringarray.MakeStringArr()
}

func pkg_forloop() {
	forloop.ForloopPractice()
}

func pkg_booksapi() {
	booksapiserver.NewMuxServer()
}
