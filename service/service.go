package service

import (
	"encoding/json"
	zendesk "github.com/MEDIGO/go-zendesk/zendesk"
	"github.com/aws/aws-sdk-go/aws"
	result "github.com/oms-services/zendesk/result"
	"net/http"
	"os"
	"time"
)

//User struct
type User struct {
	Email               string                 `json:"email,omitempty"`
	Name                string                 `json:"name,omitempty"`
	Alias               string                 `json:"alias,omitempty"`
	Details             string                 `json:"details,omitempty"`
	Moderator           bool                   `json:"moderator,omitempty"`
	Notes               string                 `json:"notes,omitempty"`
	OnlyPrivateComments bool                   `json:"onlyPrivateComments,omitempty"`
	Phone               string                 `json:"phone,omitempty"`
	RestrictedAgent     bool                   `json:"restrictedAgent,omitempty"`
	Signature           string                 `json:"signature,omitempty"`
	Suspended           bool                   `json:"suspended,omitempty"`
	TicketRestriction   string                 `json:"ticketRestriction,omitempty"`
	UserFields          map[string]interface{} `json:"userFields,omitempty"`
	Verified            bool                   `json:"verified,omitempty"`
}

//Ticket struct
type Ticket struct {
	ID               int64  `json:"ticketId,omitempty"`
	ExternalID       string `json:"externalId,omitempty"`
	Type             string `json:"type,omitempty"`
	Subject          string `json:"subject,omitempty"`
	RawSubject       string `json:"rawSubject,omitempty"`
	Description      string `json:"description,omitempty"`
	Priority         string `json:"priority,omitempty"`
	Status           string `json:"status,omitempty"`
	Recipient        string `json:"recipient,omitempty"`
	RequesterID      int64  `json:"requesterId,omitempty"`
	SubmitterID      int64  `json:"submitterId,omitempty"`
	OrganizationID   int64  `json:"organizationId,omitempty"`
	GroupID          int64  `json:"groupId,omitempty"`
	ForumTopicID     int64  `json:"forumTopicId,omitempty"`
	ProblemID        int64  `json:"problemId,omitempty"`
	DueAt            string `json:"dueAt,omitempty"`
	BrandID          int64  `json:"brandID,omitempty"`
	TicketFormID     int64  `json:"ticketFormId,omitempty"`
	FollowupSourceID int64  `json:"followupSourceId,omitempty"`
	SortBy           string `json:"sortBy,omitempty"`
	SortOrder        string `json:"sortOrder,omitempty"`
}

//Message struct
type Message struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

const layout = "2006-01-02T15:04:05"

//CreateUser Zendesk
func CreateUser(responseWriter http.ResponseWriter, request *http.Request) {

	domainName := os.Getenv("DOMAIN_NAME")
	email := os.Getenv("EMAIL") + "/token"
	apiToken := os.Getenv("API_TOKEN")

	decoder := json.NewDecoder(request.Body)
	var param User
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client, clientErr := zendesk.NewClient(domainName, email, apiToken)
	if clientErr != nil {
		result.WriteErrorResponseString(responseWriter, clientErr.Error())
		return
	}

	userDetails := zendesk.User{
		Name:                aws.String(param.Name),
		Email:               aws.String(param.Email),
		Alias:               aws.String(param.Alias),
		Details:             aws.String(param.Details),
		Moderator:           aws.Bool(param.Moderator),
		Notes:               aws.String(param.Notes),
		OnlyPrivateComments: aws.Bool(param.OnlyPrivateComments),
		Phone:               aws.String(param.Phone),
		RestrictedAgent:     aws.Bool(param.RestrictedAgent),
		Signature:           aws.String(param.Signature),
		Suspended:           aws.Bool(param.Suspended),
		TicketRestriction:   aws.String(param.TicketRestriction),
		UserFields:          aws.JSONValue(param.UserFields),
		Verified:            aws.Bool(param.Verified),
	}

	newUser, userErr := client.CreateUser(&userDetails)
	if userErr != nil {
		result.WriteErrorResponseString(responseWriter, userErr.Error())
		return
	}

	bytes, _ := json.Marshal(newUser)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//CreateTicket Zendesk
func CreateTicket(responseWriter http.ResponseWriter, request *http.Request) {

	domainName := os.Getenv("DOMAIN_NAME")
	email := os.Getenv("EMAIL") + "/token"
	apiToken := os.Getenv("API_TOKEN")

	decoder := json.NewDecoder(request.Body)
	var param Ticket
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client, clientErr := zendesk.NewClient(domainName, email, apiToken)
	if clientErr != nil {
		result.WriteErrorResponseString(responseWriter, clientErr.Error())
		return
	}

	var dueDate time.Time
	var dueErr error
	if param.DueAt != "" {
		dueDate, dueErr = time.Parse(layout, param.DueAt)
		if dueErr != nil {
			result.WriteErrorResponseString(responseWriter, dueErr.Error())
			return
		}
	}

	ticketDetails := zendesk.Ticket{
		ExternalID:  aws.String(param.ExternalID),
		Type:        aws.String(param.Type),
		Subject:     aws.String(param.Subject),
		RawSubject:  aws.String(param.RawSubject),
		Description: aws.String(param.Description),
		Priority:    aws.String(param.Priority),
		Status:      aws.String(param.Status),
		Recipient:   aws.String(param.Recipient),
		RequesterID: aws.Int64(param.RequesterID),
		DueAt:       aws.Time(dueDate),
	}

	newTicket, ticketErr := client.CreateTicket(&ticketDetails)
	if ticketErr != nil {
		result.WriteErrorResponseString(responseWriter, ticketErr.Error())
		return
	}

	bytes, _ := json.Marshal(newTicket)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//ListTicket Zendesk
func ListTicket(responseWriter http.ResponseWriter, request *http.Request) {

	domainName := os.Getenv("DOMAIN_NAME")
	email := os.Getenv("EMAIL") + "/token"
	apiToken := os.Getenv("API_TOKEN")

	decoder := json.NewDecoder(request.Body)
	var param Ticket
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client, clientErr := zendesk.NewClient(domainName, email, apiToken)
	if clientErr != nil {
		result.WriteErrorResponseString(responseWriter, clientErr.Error())
		return
	}

	listTicket := zendesk.ListOptions{
		SortBy:    param.SortBy,
		SortOrder: param.SortOrder,
	}

	ticketLists, ticketErr := client.ListTickets(&listTicket)
	if ticketErr != nil {
		result.WriteErrorResponseString(responseWriter, ticketErr.Error())
		return
	}

	bytes, _ := json.Marshal(ticketLists)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//DeleteTicket Zendesk
func DeleteTicket(responseWriter http.ResponseWriter, request *http.Request) {

	domainName := os.Getenv("DOMAIN_NAME")
	email := os.Getenv("EMAIL") + "/token"
	apiToken := os.Getenv("API_TOKEN")

	decoder := json.NewDecoder(request.Body)
	var param Ticket
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client, clientErr := zendesk.NewClient(domainName, email, apiToken)
	if clientErr != nil {
		result.WriteErrorResponseString(responseWriter, clientErr.Error())
		return
	}

	ticketErr := client.DeleteTicket(param.ID)
	if ticketErr != nil {
		result.WriteErrorResponseString(responseWriter, ticketErr.Error())
		return
	}

	message := Message{true, "Ticket deleted successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
