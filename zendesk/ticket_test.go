package zendesk

import (
	"net/http"
	"testing"
)

func TestCreateTicket(t *testing.T) {
	mockAPI := newMockAPIWithStatus(http.MethodPost, "ticket.json", http.StatusCreated)
	client := newTestClient(mockAPI)
	defer mockAPI.Close()

	out, err := client.CreateTicket(Ticket{})
	if err != nil {
		t.Fatalf("Failed to send request to create ticket field: %s", err)
	}

	expected := int64(35436)
	if out.ID != expected {
		t.Fatalf("Ticket did not have expected ID. Was: %d expected: %d", out.ID, expected)
	}
}
