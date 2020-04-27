// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A JSON object containing the Amazon Resource Name (ARN) of the gateway to
// update.
type UpdateGatewaySoftwareNowInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateGatewaySoftwareNowInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGatewaySoftwareNowInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGatewaySoftwareNowInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway that
// was updated.
type UpdateGatewaySoftwareNowOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s UpdateGatewaySoftwareNowOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateGatewaySoftwareNow = "UpdateGatewaySoftwareNow"

// UpdateGatewaySoftwareNowRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Updates the gateway virtual machine (VM) software. The request immediately
// triggers the software update.
//
// When you make this request, you get a 200 OK success response immediately.
// However, it might take some time for the update to complete. You can call
// DescribeGatewayInformation to verify the gateway is in the STATE_RUNNING
// state.
//
// A software update forces a system restart of your gateway. You can minimize
// the chance of any disruption to your applications by increasing your iSCSI
// Initiators' timeouts. For more information about increasing iSCSI Initiator
// timeouts for Windows and Linux, see Customizing Your Windows iSCSI Settings
// (https://docs.aws.amazon.com/storagegateway/latest/userguide/ConfiguringiSCSIClientInitiatorWindowsClient.html#CustomizeWindowsiSCSISettings)
// and Customizing Your Linux iSCSI Settings (https://docs.aws.amazon.com/storagegateway/latest/userguide/ConfiguringiSCSIClientInitiatorRedHatClient.html#CustomizeLinuxiSCSISettings),
// respectively.
//
//    // Example sending a request using UpdateGatewaySoftwareNowRequest.
//    req := client.UpdateGatewaySoftwareNowRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/UpdateGatewaySoftwareNow
func (c *Client) UpdateGatewaySoftwareNowRequest(input *UpdateGatewaySoftwareNowInput) UpdateGatewaySoftwareNowRequest {
	op := &aws.Operation{
		Name:       opUpdateGatewaySoftwareNow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateGatewaySoftwareNowInput{}
	}

	req := c.newRequest(op, input, &UpdateGatewaySoftwareNowOutput{})
	return UpdateGatewaySoftwareNowRequest{Request: req, Input: input, Copy: c.UpdateGatewaySoftwareNowRequest}
}

// UpdateGatewaySoftwareNowRequest is the request type for the
// UpdateGatewaySoftwareNow API operation.
type UpdateGatewaySoftwareNowRequest struct {
	*aws.Request
	Input *UpdateGatewaySoftwareNowInput
	Copy  func(*UpdateGatewaySoftwareNowInput) UpdateGatewaySoftwareNowRequest
}

// Send marshals and sends the UpdateGatewaySoftwareNow API request.
func (r UpdateGatewaySoftwareNowRequest) Send(ctx context.Context) (*UpdateGatewaySoftwareNowResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateGatewaySoftwareNowResponse{
		UpdateGatewaySoftwareNowOutput: r.Request.Data.(*UpdateGatewaySoftwareNowOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateGatewaySoftwareNowResponse is the response type for the
// UpdateGatewaySoftwareNow API operation.
type UpdateGatewaySoftwareNowResponse struct {
	*UpdateGatewaySoftwareNowOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateGatewaySoftwareNow request.
func (r *UpdateGatewaySoftwareNowResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
