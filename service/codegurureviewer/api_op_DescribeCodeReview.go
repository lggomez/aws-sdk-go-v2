// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codegurureviewer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeCodeReviewInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the code review to describe.
	//
	// CodeReviewArn is a required field
	CodeReviewArn *string `location:"uri" locationName:"CodeReviewArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCodeReviewInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCodeReviewInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCodeReviewInput"}

	if s.CodeReviewArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("CodeReviewArn"))
	}
	if s.CodeReviewArn != nil && len(*s.CodeReviewArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CodeReviewArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCodeReviewInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CodeReviewArn != nil {
		v := *s.CodeReviewArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "CodeReviewArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeCodeReviewOutput struct {
	_ struct{} `type:"structure"`

	// Information about the code review.
	CodeReview *CodeReview `type:"structure"`
}

// String returns the string representation
func (s DescribeCodeReviewOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeCodeReviewOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CodeReview != nil {
		v := s.CodeReview

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CodeReview", v, metadata)
	}
	return nil
}

const opDescribeCodeReview = "DescribeCodeReview"

// DescribeCodeReviewRequest returns a request value for making API operation for
// Amazon CodeGuru Reviewer.
//
// Returns the metadaata associated with the code review along with its status.
//
//    // Example sending a request using DescribeCodeReviewRequest.
//    req := client.DescribeCodeReviewRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguru-reviewer-2019-09-19/DescribeCodeReview
func (c *Client) DescribeCodeReviewRequest(input *DescribeCodeReviewInput) DescribeCodeReviewRequest {
	op := &aws.Operation{
		Name:       opDescribeCodeReview,
		HTTPMethod: "GET",
		HTTPPath:   "/codereviews/{CodeReviewArn}",
	}

	if input == nil {
		input = &DescribeCodeReviewInput{}
	}

	req := c.newRequest(op, input, &DescribeCodeReviewOutput{})
	return DescribeCodeReviewRequest{Request: req, Input: input, Copy: c.DescribeCodeReviewRequest}
}

// DescribeCodeReviewRequest is the request type for the
// DescribeCodeReview API operation.
type DescribeCodeReviewRequest struct {
	*aws.Request
	Input *DescribeCodeReviewInput
	Copy  func(*DescribeCodeReviewInput) DescribeCodeReviewRequest
}

// Send marshals and sends the DescribeCodeReview API request.
func (r DescribeCodeReviewRequest) Send(ctx context.Context) (*DescribeCodeReviewResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCodeReviewResponse{
		DescribeCodeReviewOutput: r.Request.Data.(*DescribeCodeReviewOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCodeReviewResponse is the response type for the
// DescribeCodeReview API operation.
type DescribeCodeReviewResponse struct {
	*DescribeCodeReviewOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCodeReview request.
func (r *DescribeCodeReviewResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
