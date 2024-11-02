package domain

import "errors"

type Feedback struct {
	CompanyName  string
	ContractType ContractType
	Plan         string
	Read         bool
	CreatedBy    string
	Content      string
}

func NewFeedback(companyName, contractType, plan, createdBy, content string) (*Feedback, error) {
	if companyName == "" {
		return nil, errors.New("company name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if createdBy == "" {
		return nil, errors.New("createby is required")
	}

	ct, err := NewContractType(contractType)
	if err != nil {
		return nil, err
	}

	return &Feedback{
		CompanyName:  companyName,
		ContractType: *ct,
		Plan:         plan,
		Read:         false,
		CreatedBy:    createdBy,
		Content:      content,
	}, nil
}
