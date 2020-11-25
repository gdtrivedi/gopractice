package errorf

import "fmt"

func ErrorfTest() {
	acctID := 10
	err := fmt.Errorf("error for Account: %d", acctID)
	fmt.Println(err.Error())
}
