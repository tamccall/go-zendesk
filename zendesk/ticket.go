package zendesk

import (
	"encoding/json"
	"time"
)

const (
	IncidentTicketType = "incident"
	ProblemTicketType  = "problem"
	QuestionTicketType = "question"
	TaskTicketType     = "task"

	UrgentTicketPriority = "urgent"
	HighTicketPriority   = "high"
	NormalTicketPriority = "normal"
	LowTicketPriority    = "low"

	NewTicketStatus     = "new"
	OpenTicketStatus    = "open"
	PendingTicketStatus = "pending"
	HoldTicketStatus    = "hold"
	SolvedTicketStatus  = "solved"
	ClosedTicketStatus  = "closed"
)

type CustomField struct {
	ID    int64  `json:"id"`
	Value string `json:"value"`
}

type SatisfactionRating struct {
	ID      int64  `json:"id"`
	Score   string `json:"score"`
	Comment string `json:"comment"`
}

type Ticket struct {
	ID              int64     `json:"id,omitempty"`
	URL             string    `json:"url,omitempty"`
	ExternalID      string    `json:"external_id,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	Type            string    `json:"type,omitempty"`
	Subject         string    `json:"subject,omitempty"`
	RawSubject      string    `json:"raw_subject,omitempty"`
	Description     string    `json:"description,omitempty"`
	Priority        string    `json:"priority,omitempty"`
	Status          string    `json:"status,omitempty"`
	Recipient       string    `json:"recipient,omitempty"`
	RequesterID     int64     `json:"requester_id,omitempty"`
	SubmitterID     int64     `json:"submitter_id, omitempty"`
	AssigneeID      int64     `json:"assignee_id, omitempty"`
	OrganizationID  int64     `json:"organization_id, omitempty"`
	GroupID         int64     `json:"group_id,omitempty"`
	CollaboratorIDs []int64   `json:"collaborator_ids, omitempty"`
	FollowerIDs     []int64   `json:"follower_ids, omitempty"`
	ProblemID       int64     `json:"problem_id, omitempty"`
	HasIncidents    bool      `json:"has_incidents,omitempty"`
	DueAt           time.Time `json:"due_at, omitempty"`
	Tags            []string  `json:"tags,omitempty"`
	//TODO: add via object https://developer.zendesk.com/rest_api/docs/support/ticket_audits.html#the-via-object
	CustomFields         []CustomField      `json:"custom_fields,omitempty"`
	SatisfactionRating   SatisfactionRating `json:"satisfaction_rating,omitempty"`
	SharingAgreementIDs []int64            `json:"sharing_agreement_ids, omitempty"`
}

type TicketAPI interface {
	CreateTicket(ticket Ticket) (Ticket, error)
}

func (z *Client) CreateTicket(ticket Ticket) (Ticket, error) {
	var data, result struct {
		Ticket Ticket `json:"ticket"`
	}
	data.Ticket = ticket

	body, err := z.Post("/tickets.json", data)
	if err != nil {
		return Ticket{}, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return Ticket{}, err
	}
	return result.Ticket, nil
}
