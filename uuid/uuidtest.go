package uuid

import (
	"fmt"

	"github.com/google/uuid"
)

func UUIDTest() {
	fmt.Println("UUID: ", uuid.New())
}
