package main

import (
	"encoding/json"
	"os"
)


type Policy struct {
	PolicyName     string
	PolicyDocument PolicyDocument
}

type PolicyDocument struct {
	Version    string
	Statement  []Statement
}

type Statement struct {
	Sid      string
	Effect   string
	Action   []string
	Resource string
}

func VerifyPolicy(filePath string) (bool, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	var policy Policy
	if err := json.Unmarshal(data, &policy); err != nil {
		return false, err
	}

	for _, statement := range policy.PolicyDocument.Statement {
		if statement.Resource == "*" {
			return false, nil
		}
	}

	return true, nil
}