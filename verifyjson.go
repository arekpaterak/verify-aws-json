package main

import (
	"encoding/json"
	"os"
)

type PolicyDocument struct {
	Version   string
	Statement []struct {
		Sid     string
		Effect  string
		Action  []string
		Resource string
	}
}

type Policy struct {
	PolicyName     string
	PolicyDocument PolicyDocument
}

func verifyPolicy(filePath string) (bool, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	var policy Policy
	err = json.Unmarshal(data, &policy)
	if err != nil {
		return false, err
	}

	for _, statement := range policy.PolicyDocument.Statement {
		if statement.Resource == "*" {
			return false, nil
		}
	}

	return true, nil
}