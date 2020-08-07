package stringarray

import (
	"strings"
)

//type strArr []string
type StrArr []string
type str2DArr [][]string

func stringArrayTests() {
	// START: split test
	//inputStrArr := make([]string, 0, 10) //[10]string{"a","b","c","d","e","f","g","h","i","j"}
	//inputStrArr = append(append(append(append(append(append(append(append(append(append(inputStrArr, "a"), "b"), "c"), "d"), "e"), "f"), "g"), "h"), "i"), "j")
	//inputStrArr = inputStrArr[:0] // clear the slice
	//inputStrArr = append(append(append(inputStrArr, "a"), "b"), "c")
	//splitArr := split(inputStrArr, 4)
	//fmt.Println("After Split: ", splitArr)
	// END: split test

	// START: split1 test
	//inputStrArr1 := strArr{""}
	//elements := split1(inputStrArr1, 5)
	//fmt.Println("After Split: ", elements)
	//fmt.Println("len(elements) : ", len(elements))
	//fmt.Println("len(elements[0]) : ", len(elements[0]))
	//fmt.Println("len(elements[0][0]) : ", len(elements[0][0]))
	// END: split1 test
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

func split1(strings StrArr, chunkSize int) str2DArr {
	var chunk StrArr

	chunks := make(str2DArr, 0, len(strings)/chunkSize+1)

	for len(strings) >= chunkSize {
		chunk, strings = strings[:chunkSize], strings[chunkSize:]
		chunks = append(chunks, chunk)
	}

	if len(strings) > 0 {
		chunks = append(chunks, strings)
	}

	return chunks
}

// Split splits given list of string and chunk size, this function splits the string array into multi dimension array of fixed lim size. Returns empty [] if input strings [] is empty.
func Split(strings []string, chunkSize int) [][]string {
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

// Removes elements from string array/slice which matches to selector string and returns result array/slice.
// This functon also accepts a bool parameter to trim the array/slice element before matching.
func RemoveElements(inputStrArr StrArr, selector string, trimElementBeforeMatch bool) StrArr {
	var outputStrArr StrArr
	for _, str := range inputStrArr {
		strToCompare := str
		if trimElementBeforeMatch {
			strToCompare = strings.TrimSpace(strToCompare)
		}

		if strToCompare != selector {
			outputStrArr = append(outputStrArr, strToCompare)
		}
	}
	return outputStrArr
}
