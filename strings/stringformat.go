package strings

import (
	"fmt"
)

func FormatString() {
	baseUrl := "https://api.hubapi.com"
	workflowId := 99999
	email := "email@example.com"
	url := fmt.Sprintf("%s/automation/v2/workflows/%d/enrollments/contacts/%s", baseUrl, workflowId, email)
	fmt.Println("URL: ", url)
}