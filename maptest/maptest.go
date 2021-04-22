package maptest

import "fmt"

func MapTest() {
	testMap()
}

func testMap() {
	m := make(map[string]bool)
	m["a"] = false
	m["mu"] = true
	optionName := "mu1"
	_, ok := m[string(optionName)]
	fmt.Print("ok = ", ok)
}
