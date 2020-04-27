// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeCachediSCSIVolumesInput struct {
	_ struct{} `type:"structure"`

	// An array of strings where each string represents the Amazon Resource Name
	// (ARN) of a cached volume. All of the specified cached volumes must be from
	// the same gateway. Use ListVolumes to get volume ARNs for a gateway.
	//
	// VolumeARNs is a required field
	VolumeARNs []string `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeCachediSCSIVolumesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCachediSCSIVolumesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCachediSCSIVolumesInput"}

	if s.VolumeARNs == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeARNs"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the following fields:
type DescribeCachediSCSIVolumesOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects where each object contains metadata about one cached
	// volume.
	CachediSCSIVolumes []CachediSCSIVolume `type:"list"`
}

// String returns the string representation
func (s DescribeCachediSCSIVolumesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeCachediSCSIVolumes = "DescribeCachediSCSIVolumes"

// DescribeCachediSCSIVolumesRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Returns a description of the gateway volumes specified in the request. This
// operation is only supported in the cached volume gateway types.
//
// The list of gateway volumes in the request must be from one gateway. In the
// response, AWS Storage Gateway returns volume information sorted by volume
// Amazon Resource Name (ARN).
//
//    // Example sending a request using DescribeCachediSCSIVolumesRequest.
//    req := client.DescribeCachediSCSIVolumesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/DescribeCachediSCSIVolumes
func (c *Client) DescribeCachediSCSIVolumesRequest(input *DescribeCachediSCSIVolumesInput) DescribeCachediSCSIVolumesRequest {
	op := &aws.Operation{
		Name:       opDescribeCachediSCSIVolumes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCachediSCSIVolumesInput{}
	}

	req := c.newRequest(op, input, &DescribeCachediSCSIVolumesOutput{})
	return DescribeCachediSCSIVolumesRequest{Request: req, Input: input, Copy: c.DescribeCachediSCSIVolumesRequest}
}

// DescribeCachediSCSIVolumesRequest is the request type for the
// DescribeCachediSCSIVolumes API operation.
type DescribeCachediSCSIVolumesRequest struct {
	*aws.Request
	Input *DescribeCachediSCSIVolumesInput
	Copy  func(*DescribeCachediSCSIVolumesInput) DescribeCachediSCSIVolumesRequest
}

// Send marshals and sends the DescribeCachediSCSIVolumes API request.
func (r DescribeCachediSCSIVolumesRequest) Send(ctx context.Context) (*DescribeCachediSCSIVolumesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeCachediSCSIVolumesResponse{
		DescribeCachediSCSIVolumesOutput: r.Request.Data.(*DescribeCachediSCSIVolumesOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeCachediSCSIVolumesResponse is the response type for the
// DescribeCachediSCSIVolumes API operation.
type DescribeCachediSCSIVolumesResponse struct {
	*DescribeCachediSCSIVolumesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeCachediSCSIVolumes request.
func (r *DescribeCachediSCSIVolumesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
