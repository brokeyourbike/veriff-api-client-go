package veriff

import (
	"context"
	"fmt"
	"net/http"
)

type CreateSessionPayload struct {
	Verification struct {
		Person struct {
			FirstName   string `json:"firstName,omitempty"`
			LastName    string `json:"lastName,omitempty"`
			DateOfBirth string `json:"dateOfBirth,omitempty"`
			Gender      string `json:"gender,omitempty"`
		} `json:"person"`
		Callback   string `json:"callback,omitempty"`
		VendorData string `json:"vendorData"`
	} `json:"verification"`
}

type CreateSessionResponse struct {
	Status       string `json:"status"`
	Verification struct {
		ID           string `json:"id"`
		URL          string `json:"url"`
		VendorData   string `json:"vendorData"`
		Host         string `json:"host"`
		Status       string `json:"status"`
		SessionToken string `json:"sessionToken"`
	} `json:"verification"`
}

func (c *client) CreateSession(ctx context.Context, payload CreateSessionPayload) (data CreateSessionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodPost, "/v1/sessions", payload)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}
