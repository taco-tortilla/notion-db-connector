package request

type FeedbackRequest struct {
	CompanyName  string `json:"company_name"`
	ContractType string `json:"contract_type"`
	Plan         string `json:"plan"`
	CreatedBy    string `json:"created_by"`
	Body         string `json:"body"`
}
