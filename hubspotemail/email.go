package hubspotemail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/bold-commerce/go-hubspot/hubspot"
)

var (
	baseUrl = "https://api.hubapi.com"
)

type Client struct {
	baseUrl string
	apiKey  string
}

type HubspotContact struct {
	Properties	[]HubspotContactProperty `json:"properties"`
}

type HubspotContactProperty struct {
	Property	string `json:"property"`
	Value 		string `json:"value"`
}

func NewClient(baseUrl, apiKey string) *Client {
	return &Client{
		baseUrl: baseUrl,
		apiKey:  apiKey,
	}
}

type HubspotAutoRequest struct {
	GrantType		string `json:"grant_type`
	ClientID		string `json:"client_id"`
	ClientSecret	string `json:"client_secret"`
	RefreshToken	string `json:"refresh_token"`
}

type HubspotAutoResponse struct {
	AccessToken		string 	`json:"access_token"`
	RefreshToken	string 	`json:"refresh_token"`
	ExpiresIn		int		`json:"expires_in"`
	Status			string	`json:"status"`
	Message			string	`json:"message"`
	CorrelationID	string	`json:"correlationId"`
	RequestID		string	`json:"requestId"`
	/*
	"status": "EXAMPLE_ERROR_CODE",
	  "message": "A human readable error message",
	  "correlationId": "da1e1d2f-fa6b-472a-be7b-7ab9b9605d59",
	  "requestId": "ecc0b50659814f7ca37f5d49cdf9cbd3"
	 */
}

func (c *Client) AuthToken() {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("client_id", "client_id_from_secrets_manager")
	data.Set("client_secret", "client_secret_from_secrets_manager")
	data.Set("refresh_token", "refresh_token_from_secrets_manager")

	url := "https://api.hubapi.com/oauth/v1/token" //call with Bearer token in request header.

	httpClient := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error2: ", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	log.Printf("GetAccessToken: Calling hubspot API to get access token")
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Error3: ", err)
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error4: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error5: ", resBody)
	}

	var authRes HubspotAutoResponse
	if err := json.Unmarshal(resBody, &authRes); err != nil {
		fmt.Println("Error6: ", err)
	}

	fmt.Println("AuthResponse: ", authRes)
	fmt.Println("AccessToken: ", authRes.AccessToken)
}

func SendEmail() {
	client := hubspot.NewClient("https://api.hubapi.com", "api-key")

	emailId := 33031315159
	err := client.SingleEmail(emailId, "gautam.trivedi@gmail.com")
	if err != nil {
		log.Fatalf("hubspot error: %s", err.Error())
	}
}

// Create hubspot contact if not present.
func (c *Client) CreateHubspotContact(email string) error {
	contactProperties := make([]HubspotContactProperty, 0)
	contactPropertyEmail := &HubspotContactProperty{
		// At least one property is required from all allowed.
		// Refer https://legacydocs.hubspot.com/docs/methods/contacts/create_contact for more details.
		Property: "email",
		Value:    email,
	}

	contactPropertyFirstname := &HubspotContactProperty{
		// At least one property is required from all allowed.
		// Refer https://legacydocs.hubspot.com/docs/methods/contacts/create_contact for more details.
		Property: "firstname",
		Value:    "Gautam",
	}

	contactPropertyLastname := &HubspotContactProperty{
		// At least one property is required from all allowed.
		// Refer https://legacydocs.hubspot.com/docs/methods/contacts/create_contact for more details.
		Property: "lastname",
		Value:    "Trivedi",
	}

	contactPropertyInstall := &HubspotContactProperty{
		// At least one property is required from all allowed.
		// Refer https://legacydocs.hubspot.com/docs/methods/contacts/create_contact for more details.
		Property: "install_name",
		Value:    "Install-123",
	}

	contactPropertyAccount := &HubspotContactProperty{
		// At least one property is required from all allowed.
		// Refer https://legacydocs.hubspot.com/docs/methods/contacts/create_contact for more details.
		Property: "account_name",
		Value:    "Account-123",
	}

	contactPropertyMessage := &HubspotContactProperty{
		// At least one property is required from all allowed.
		// Refer https://legacydocs.hubspot.com/docs/methods/contacts/create_contact for more details.
		Property: "message",
		Value:    "Test message goes here. Blah blah blah",
	}
	contactProperties = append(append(append(append(append(append(contactProperties, *contactPropertyEmail), *contactPropertyFirstname), *contactPropertyLastname), *contactPropertyInstall), *contactPropertyAccount), *contactPropertyMessage)
	hubspotContact := HubspotContact{Properties: contactProperties}

	reqBody, err := json.Marshal(hubspotContact)
	if err != nil {
		return fmt.Errorf("CreateHubspotContact: invalid request: %s", err.Error())
	}

	url := fmt.Sprintf("%s/contacts/v1/contact/createOrUpdate/email/%s", baseUrl, email) //call with Bearer token in request header.

	httpClient := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return errors.Wrap(err, "CreateHubspotContact: Unable to create request.")
	}

	req.Header.Set("Content-Type", "application/json")
	token := "grabbed token from secrets" //TODO: replace with actual token from google secrets.
	req.Header.Set("Authorization","Bearer " + token)

	log.Printf("CreateHubspotContact: Calling hubspot API to crate contact")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Errorf("CreateHubspotContact: Failed to create contact.", err)
		return errors.Wrap(err, "CreateHubspotContact: Failed to create contact.")
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "CreateHubspotContact: Failed to read response.")
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusConflict {
		return errors.Wrap(fmt.Errorf("CreateHubspotContact: Error: %s details: %s\n", resp.Status, resBody), "Non OK response received.")
	}

	return nil
}

func (c *Client) EnrollContactToWorkflow(email string) error {
	workflowId := 1234567890 //TODO: replade with actual workflow id.
	url := fmt.Sprintf("%s/automation/v2/workflows/%d/enrollments/contacts/%s", c.baseUrl, workflowId, email)

	httpClient := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return errors.Wrap(err, "EnrollContactToWorkflow: Unable to create request.")
	}

	req.Header.Set("Content-Type", "application/json")
	token := "grabbed token from secrets" //TODO: replace with actual token from google secrets.
	req.Header.Set("Authorization","Bearer " + token)

	log.Printf("EnrollContactToWorkflow: Calling hubspot API to enroll contact")
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Errorf("CreateHubspotContact: Failed to enroll contact.", err)
		return errors.Wrap(err, "EnrollContactToWorkflow: Failed to enroll contact.")
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "EnrollContactToWorkflow: Failed to read response.")
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return errors.Wrap(fmt.Errorf("EnrollContactToWorkflow: Error: %s details: %s\n", resp.Status, resBody), "Non OK response received.")
	}

	return nil
}