// Package acm is used to interact with the AWS Certificate Manager API.
package acm

import (
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
	"github.com/aws/smithy-go/middleware"
)

// ClientInterface for interacting with CloudWatch.
type ClientInterface interface {
	DescribeCertificate(params *acm.DescribeCertificateInput, optFns ...func(*acm.Options)) (*acm.DescribeCertificateOutput, error)
}

// MockClient used for testing purposes.
type MockClient struct {
	Certificate    types.CertificateDetail
	ResultMetadata middleware.Metadata
}

// DescribeCertificate mocks the CloudWatch API.
func (m *MockClient) DescribeCertificate(params *acm.DescribeCertificateInput, optFns ...func(*acm.Options)) (*acm.DescribeCertificateOutput, error) {
	return &acm.DescribeCertificateOutput{
		Certificate:    &m.Certificate,
		ResultMetadata: m.ResultMetadata,
	}, nil
}
