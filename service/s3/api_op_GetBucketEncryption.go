// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the default encryption configuration for an Amazon S3 bucket. For
// information about the Amazon S3 default encryption feature, see Amazon S3
// Default Bucket Encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html).  <p>
// To use this operation, you must have permission to perform the
// <code>s3:GetEncryptionConfiguration</code> action. The bucket owner has this
// permission by default. The bucket owner can grant this permission to others. For
// more information about permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources">Permissions
// Related to Bucket Subresource Operations</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html">Managing
// Access Permissions to Your Amazon S3 Resources</a>.</p> <p>The following
// operations are related to <code>GetBucketEncryption</code>:</p> <ul> <li> <p>
// <a>PutBucketEncryption</a> </p> </li> <li> <p> <a>DeleteBucketEncryption</a>
// </p> </li> </ul>
func (c *Client) GetBucketEncryption(ctx context.Context, params *GetBucketEncryptionInput, optFns ...func(*Options)) (*GetBucketEncryptionOutput, error) {
	stack := middleware.NewStack("GetBucketEncryption", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpGetBucketEncryptionMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetBucketEncryptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketEncryption(options.Region), middleware.Before)
	addUpdateEndpointMiddleware(stack, options)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "GetBucketEncryption",
			Err:           err,
		}
	}
	out := result.(*GetBucketEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketEncryptionInput struct {
	// The name of the bucket from which the server-side encryption configuration is
	// retrieved.
	Bucket *string
}

type GetBucketEncryptionOutput struct {
	// Specifies the default server-side-encryption configuration.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpGetBucketEncryptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpGetBucketEncryption{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketEncryption{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetBucketEncryption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketEncryption",
	}
}
