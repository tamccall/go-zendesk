package zendesk

import "time"

type CustomField struct {
	ID    int64  `json:"id"`
	Value string `json:"value"`
}

type SatisfactionRating struct {
	ID      int64  `json:"id"`
	Score   string `json:"score"`
	Comment string `json:"comment"`
}

//{
//"id":               35436,
//"url":              "https://company.zendesk.com/api/v2/tickets/35436.json",
//"external_id":      "ahg35h3jh",
//"created_at":       "2009-07-20T22:55:29Z",
//"updated_at":       "2011-05-05T10:38:52Z",
//"type":             "incident",
//"subject":          "Help, my printer is on fire!",
//"raw_subject":      "{{dc.printer_on_fire}}",
//"description":      "The fire is very colorful.",
//"priority":         "high",
//"status":           "open",
//"recipient":        "support@company.com",
//"requester_id":     20978392,
//"submitter_id":     76872,
//"assignee_id":      235323,
//"organization_id":  509974,
//"group_id":         98738,
//"collaborator_ids": [35334, 234],
//"follower_ids":     [35334, 234], // This functionally is the same as collaborators for now.
//"problem_id":       9873764,
//"has_incidents":    false,
//"due_at":           null,
//"tags":             ["enterprise", "other_tag"],
//"via": {
//"channel": "web"
//},
//"custom_fields": [
//{
//"id":    27642,
//"value": "745"
//},
//{
//"id":    27648,
//"value": "yes"
//}
//],
//"satisfaction_rating": {
//"id": 1234,
//"score": "good",
//"comment": "Great support!"
//},
//"sharing_agreement_ids": [84432]
//}

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
	SharingAggreementIDs []int64            `json:"sharing_aggreement_ids, omitempty"`
}
