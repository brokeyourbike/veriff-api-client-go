package veriff

import (
	"context"
	"fmt"
	"io"
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
	req, err := c.newRequest(ctx, http.MethodPost, "/v1/sessions", payload, "")
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type SessionDecisionResponse struct {
	Status       string `json:"status"`
	Verification struct {
		ID     string `json:"id"`
		Code   int    `json:"code"`
		Status string `json:"status"`
		Person struct {
			FirstName   string `json:"firstName"`
			LastName    string `json:"lastName"`
			DateOfBirth Time   `json:"dateOfBirth"`
		} `json:"person"`
		Document struct {
			Type                string `json:"type"`
			Number              string `json:"number"`
			Country             string `json:"country"`
			Remarks             string `json:"remarks"`
			State               string `json:"state"`
			PlaceOfIssue        string `json:"placeOfIssue"`
			ValidUntil          Time   `json:"validUntil"`
			FirstIssue          Time   `json:"firstIssue"`
			IssueNumber         string `json:"issueNumber"`
			IssuedBy            string `json:"issuedBy"`
			NfcValidated        string `json:"nfcValidated"`
			ResidencePermitType string `json:"residencePermitType"`
			PortraitIsVisible   string `json:"portraitIsVisible"`
			SignatureIsVisible  string `json:"signatureIsVisible"`
		} `json:"document"`
	} `json:"verification"`
}

func (c *client) SessionDecision(ctx context.Context, sessionID string) (data SessionDecisionResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/v1/sessions/%s/decision", sessionID), nil, sessionID)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

type SessionMediaResponse struct {
	Status string `json:"status"`
	Videos []struct {
		ID        string  `json:"id"`
		SessionID string  `json:"sessionId"`
		Context   string  `json:"context"`
		Duration  float64 `json:"duration"`
		Mimetype  string  `json:"mimetype"`
		Name      string  `json:"name"`
		Size      int64   `json:"size"`
		URL       string  `json:"url"`
	} `json:"videos"`
	Images []struct {
		ID        string `json:"id"`
		SessionID string `json:"sessionId"`
		Context   string `json:"context"`
		Mimetype  string `json:"mimetype"`
		Name      string `json:"name"`
		URL       string `json:"url"`
		Size      int64  `json:"size"`
	} `json:"images"`
}

func (c *client) SessionMedia(ctx context.Context, sessionID string) (data SessionMediaResponse, err error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/v1/sessions/%s/media", sessionID), nil, sessionID)
	if err != nil {
		return data, fmt.Errorf("failed to create request: %w", err)
	}

	req.DecodeTo(&data)
	return data, c.do(ctx, req)
}

func (c *client) DownloadMedia(ctx context.Context, mediaID string, dst io.Writer) (err error) {
	req, err := c.newRequest(ctx, http.MethodGet, fmt.Sprintf("/v1/media/%s", mediaID), nil, mediaID)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.PipeTo(dst)
	return c.do(ctx, req)
}
