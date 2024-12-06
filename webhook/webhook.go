package webhook

import "github.com/brokeyourbike/veriff-api-client-go"

type IncomingVerificationEvent struct {
	Status     string `json:"status" validate:"required"`
	Code       int    `json:"code" validate:"required"`
	Action     string `json:"action" validate:"required"`
	VendorData string `json:"vendorData" validate:"required"`
}

type IncomingDecisionEvent struct {
	Status       string `json:"status" validate:"required"`
	Verification struct {
		ID         string `json:"id" validate:"required"`
		Code       int    `json:"code" validate:"required"`
		Status     string `json:"status" validate:"required"`
		VendorData string `json:"vendorData"`
		Document   struct {
			Type       string       `json:"type"`
			Number     string       `json:"number"`
			Country    string       `json:"country"`
			ValidUntil *veriff.Time `json:"validUntil"`
		} `json:"document"`
	} `json:"verification" validate:"required"`
}
