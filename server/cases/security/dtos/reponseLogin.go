package dtos

type ResponseLogin struct {
	RequestMfa bool   `json:"request_mfa"`
	Token      string `json:"token"`
}
