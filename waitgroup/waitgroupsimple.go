package waitgroup

import (
	"fmt"
	"sync"
)

func WaitgroupSimpleTest() {
	fmt.Println("Hello World")

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go myFunc(&waitgroup)
	waitgroup.Wait()

	fmt.Println("Finished Execution")
}

func myFunc(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my goroutine")
	waitgroup.Done()
}
