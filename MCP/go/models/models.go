package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// AcademicCertificateSchema represents the AcademicCertificateSchema schema from the OpenAPI specification
type AcademicCertificateSchema struct {
	Certificatedata map[string]interface{} `json:"CertificateData"`
	Issuedby map[string]interface{} `json:"IssuedBy"`
	Validfromdate string `json:"validFromDate"`
	Issuedat string `json:"issuedAt"`
	Language string `json:"language"`
	Number int `json:"number"`
	Status string `json:"status"`
	Issuedto map[string]interface{} `json:"IssuedTo"`
	Issuedate string `json:"issueDate"`
	TypeField string `json:"type"`
	Name string `json:"name"`
}

// ConsentArtifactSchema represents the ConsentArtifactSchema schema from the OpenAPI specification
type ConsentArtifactSchema struct {
	Consent map[string]interface{} `json:"consent"`
	Signature map[string]interface{} `json:"signature"`
}
