package domain

import "errors"

type ContractType struct {
	Value string
}

func NewContractType(value string) (*ContractType, error) {
	validTypes := map[string]bool{
		"contract": true,
		"trial":    true,
	}

	if !validTypes[value] {
		return nil, errors.New("invalid contract type")
	}

	return &ContractType{Value: value}, nil
}
