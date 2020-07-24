package main

import "fmt"

func stringArrayTests() {
	// START: split test
	strArr := make([]string, 0, 10) //[10]string{"a","b","c","d","e","f","g","h","i","j"}
	strArr = append(append(append(append(append(append(append(append(append(append(strArr, "a"), "b"), "c"), "d"), "e"), "f"), "g"), "h"), "i"), "j")
	strArr = strArr[:0] // clear the slice
	strArr = append(append(append(strArr, "a"), "b"), "c")
	splitArr := split(strArr, 4)
	fmt.Println("After Split: ", splitArr)
	// END: split test
}

// split given list of string and chunk size, this function splits the string array into multi dimension array of fixed lim size
func split(strings []string, chunkSize int) [][]string {
	var chunk []string

	chunks := make([][]string, 0, len(strings)/chunkSize+1)

	for len(strings) >= chunkSize {
		chunk, strings = strings[:chunkSize], strings[chunkSize:]
		chunks = append(chunks, chunk)
	}

	if len(strings) > 0 {
		chunks = append(chunks, strings)
	}

	return chunks
}
