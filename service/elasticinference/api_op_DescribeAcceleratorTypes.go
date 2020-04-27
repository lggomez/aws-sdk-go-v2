// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticinference

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeAcceleratorTypesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeAcceleratorTypesInput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAcceleratorTypesInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	return nil
}

type DescribeAcceleratorTypesOutput struct {
	_ struct{} `type:"structure"`

	// The available accelerator types.
	AcceleratorTypes []AcceleratorType `locationName:"acceleratorTypes" type:"list"`
}

// String returns the string representation
func (s DescribeAcceleratorTypesOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeAcceleratorTypesOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AcceleratorTypes != nil {
		v := s.AcceleratorTypes

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "acceleratorTypes", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDescribeAcceleratorTypes = "DescribeAcceleratorTypes"

// DescribeAcceleratorTypesRequest returns a request value for making API operation for
// Amazon Elastic  Inference.
//
// Describes the accelerator types available in a given region, as well as their
// characteristics, such as memory and throughput.
//
//    // Example sending a request using DescribeAcceleratorTypesRequest.
//    req := client.DescribeAcceleratorTypesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elastic-inference-2017-07-25/DescribeAcceleratorTypes
func (c *Client) DescribeAcceleratorTypesRequest(input *DescribeAcceleratorTypesInput) DescribeAcceleratorTypesRequest {
	op := &aws.Operation{
		Name:       opDescribeAcceleratorTypes,
		HTTPMethod: "GET",
		HTTPPath:   "/describe-accelerator-types",
	}

	if input == nil {
		input = &DescribeAcceleratorTypesInput{}
	}

	req := c.newRequest(op, input, &DescribeAcceleratorTypesOutput{})
	return DescribeAcceleratorTypesRequest{Request: req, Input: input, Copy: c.DescribeAcceleratorTypesRequest}
}

// DescribeAcceleratorTypesRequest is the request type for the
// DescribeAcceleratorTypes API operation.
type DescribeAcceleratorTypesRequest struct {
	*aws.Request
	Input *DescribeAcceleratorTypesInput
	Copy  func(*DescribeAcceleratorTypesInput) DescribeAcceleratorTypesRequest
}

// Send marshals and sends the DescribeAcceleratorTypes API request.
func (r DescribeAcceleratorTypesRequest) Send(ctx context.Context) (*DescribeAcceleratorTypesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAcceleratorTypesResponse{
		DescribeAcceleratorTypesOutput: r.Request.Data.(*DescribeAcceleratorTypesOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAcceleratorTypesResponse is the response type for the
// DescribeAcceleratorTypes API operation.
type DescribeAcceleratorTypesResponse struct {
	*DescribeAcceleratorTypesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAcceleratorTypes request.
func (r *DescribeAcceleratorTypesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
