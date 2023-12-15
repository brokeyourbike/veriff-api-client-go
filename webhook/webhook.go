package webhook

type IncomingVerificationEvent struct {
	Status     string `json:"status" validate:"required"`
	Code       int    `json:"code" validate:"required"`
	Action     string `json:"action" validate:"required"`
	VendorData string `json:"vendorData" validate:"required"`
}

type IncomingDecision struct {
	Status       string `json:"status" validate:"required"`
	Verification struct {
		ID         string `json:"id" validate:"required"`
		Code       int    `json:"code" validate:"required"`
		Status     string `json:"status" validate:"required"`
		VendorData string `json:"vendorData"`
	} `json:"verification" validate:"required"`
}
