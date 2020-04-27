// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateCachediSCSIVolumeInput struct {
	_ struct{} `type:"structure"`

	// A unique identifier that you use to retry a request. If you retry a request,
	// use the same ClientToken you specified in the initial request.
	//
	// ClientToken is a required field
	ClientToken *string `min:"5" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// True to use Amazon S3 server-side encryption with your own AWS KMS key, or
	// false to use a key managed by Amazon S3. Optional.
	KMSEncrypted *bool `type:"boolean"`

	// The Amazon Resource Name (ARN) of the AWS KMS key used for Amazon S3 server-side
	// encryption. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string `min:"7" type:"string"`

	// The network interface of the gateway on which to expose the iSCSI target.
	// Only IPv4 addresses are accepted. Use DescribeGatewayInformation to get a
	// list of the network interfaces available on a gateway.
	//
	// Valid Values: A valid IP address.
	//
	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `type:"string" required:"true"`

	// The snapshot ID (e.g. "snap-1122aabb") of the snapshot to restore as the
	// new cached volume. Specify this field if you want to create the iSCSI storage
	// volume from a snapshot otherwise do not include this field. To list snapshots
	// for your account use DescribeSnapshots (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshots.html)
	// in the Amazon Elastic Compute Cloud API Reference.
	SnapshotId *string `type:"string"`

	// The ARN for an existing volume. Specifying this ARN makes the new volume
	// into an exact copy of the specified existing volume's latest recovery point.
	// The VolumeSizeInBytes value for this new volume must be equal to or larger
	// than the size of the existing volume, in bytes.
	SourceVolumeARN *string `min:"50" type:"string"`

	// A list of up to 50 tags that you can assign to a cached volume. Each tag
	// is a key-value pair.
	//
	// Valid characters for key and value are letters, spaces, and numbers that
	// you can represent in UTF-8 format, and the following special characters:
	// + - = . _ : / @. The maximum length of a tag's key is 128 characters, and
	// the maximum length for a tag's value is 256 characters.
	Tags []Tag `type:"list"`

	// The name of the iSCSI target used by an initiator to connect to a volume
	// and used as a suffix for the target ARN. For example, specifying TargetName
	// as myvolume results in the target ARN of arn:aws:storagegateway:us-east-2:111122223333:gateway/sgw-12A3456B/target/iqn.1997-05.com.amazon:myvolume.
	// The target name must be unique across all volumes on a gateway.
	//
	// If you don't specify a value, Storage Gateway uses the value that was previously
	// used for this volume as the new target name.
	//
	// TargetName is a required field
	TargetName *string `min:"1" type:"string" required:"true"`

	// The size of the volume in bytes.
	//
	// VolumeSizeInBytes is a required field
	VolumeSizeInBytes *int64 `type:"long" required:"true"`
}

// String returns the string representation
func (s CreateCachediSCSIVolumeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCachediSCSIVolumeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCachediSCSIVolumeInput"}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 5 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 5))
	}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}
	if s.KMSKey != nil && len(*s.KMSKey) < 7 {
		invalidParams.Add(aws.NewErrParamMinLen("KMSKey", 7))
	}

	if s.NetworkInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkInterfaceId"))
	}
	if s.SourceVolumeARN != nil && len(*s.SourceVolumeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("SourceVolumeARN", 50))
	}

	if s.TargetName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetName"))
	}
	if s.TargetName != nil && len(*s.TargetName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetName", 1))
	}

	if s.VolumeSizeInBytes == nil {
		invalidParams.Add(aws.NewErrParamRequired("VolumeSizeInBytes"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateCachediSCSIVolumeOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the volume target, which includes the iSCSI
	// name that initiators can use to connect to the target.
	TargetARN *string `min:"50" type:"string"`

	// The Amazon Resource Name (ARN) of the configured volume.
	VolumeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s CreateCachediSCSIVolumeOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateCachediSCSIVolume = "CreateCachediSCSIVolume"

// CreateCachediSCSIVolumeRequest returns a request value for making API operation for
// AWS Storage Gateway.
//
// Creates a cached volume on a specified cached volume gateway. This operation
// is only supported in the cached volume gateway type.
//
// Cache storage must be allocated to the gateway before you can create a cached
// volume. Use the AddCache operation to add cache storage to a gateway.
//
// In the request, you must specify the gateway, size of the volume in bytes,
// the iSCSI target name, an IP address on which to expose the target, and a
// unique client token. In response, the gateway creates the volume and returns
// information about it. This information includes the volume Amazon Resource
// Name (ARN), its size, and the iSCSI target ARN that initiators can use to
// connect to the volume target.
//
// Optionally, you can provide the ARN for an existing volume as the SourceVolumeARN
// for this cached volume, which creates an exact copy of the existing volume’s
// latest recovery point. The VolumeSizeInBytes value must be equal to or larger
// than the size of the copied volume, in bytes.
//
//    // Example sending a request using CreateCachediSCSIVolumeRequest.
//    req := client.CreateCachediSCSIVolumeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/storagegateway-2013-06-30/CreateCachediSCSIVolume
func (c *Client) CreateCachediSCSIVolumeRequest(input *CreateCachediSCSIVolumeInput) CreateCachediSCSIVolumeRequest {
	op := &aws.Operation{
		Name:       opCreateCachediSCSIVolume,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCachediSCSIVolumeInput{}
	}

	req := c.newRequest(op, input, &CreateCachediSCSIVolumeOutput{})
	return CreateCachediSCSIVolumeRequest{Request: req, Input: input, Copy: c.CreateCachediSCSIVolumeRequest}
}

// CreateCachediSCSIVolumeRequest is the request type for the
// CreateCachediSCSIVolume API operation.
type CreateCachediSCSIVolumeRequest struct {
	*aws.Request
	Input *CreateCachediSCSIVolumeInput
	Copy  func(*CreateCachediSCSIVolumeInput) CreateCachediSCSIVolumeRequest
}

// Send marshals and sends the CreateCachediSCSIVolume API request.
func (r CreateCachediSCSIVolumeRequest) Send(ctx context.Context) (*CreateCachediSCSIVolumeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateCachediSCSIVolumeResponse{
		CreateCachediSCSIVolumeOutput: r.Request.Data.(*CreateCachediSCSIVolumeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateCachediSCSIVolumeResponse is the response type for the
// CreateCachediSCSIVolume API operation.
type CreateCachediSCSIVolumeResponse struct {
	*CreateCachediSCSIVolumeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateCachediSCSIVolume request.
func (r *CreateCachediSCSIVolumeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
