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

// Sets an analytics configuration for the bucket (specified by the analytics
// configuration ID). You can have up to 1,000 analytics configurations per bucket.
// <p>You can choose to have storage class analysis export analysis reports sent to
// a comma-separated values (CSV) flat file. See the <code>DataExport</code>
// request element. Reports are updated daily and are based on the object filters
// that you configure. When selecting data export, you specify a destination bucket
// and an optional destination prefix where the file is written. You can export the
// data to a destination bucket in a different account. However, the destination
// bucket must be in the same Region as the bucket that you are making the PUT
// analytics configuration to. For more information, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html">Amazon
// S3 Analytics – Storage Class Analysis</a>. </p> <important> <p>You must create a
// bucket policy on the destination bucket where the exported file is written to
// grant permissions to Amazon S3 to write objects to the bucket. For an example
// policy, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html#example-bucket-policies-use-case-9">Granting
// Permissions for Amazon S3 Inventory and Storage Class Analysis</a>.</p>
// </important> <p>To use this operation, you must have permissions to perform the
// <code>s3:PutAnalyticsConfiguration</code> action. The bucket owner has this
// permission by default. The bucket owner can grant this permission to others. For
// more information about permissions, see <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources">Permissions
// Related to Bucket Subresource Operations</a> and <a
// href="https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html">Managing
// Access Permissions to Your Amazon S3 Resources</a>.</p> <p class="title">
// <b>Special Errors</b> </p> <ul> <li> <ul> <li> <p> <i>HTTP Error: HTTP 400 Bad
// Request</i> </p> </li> <li> <p> <i>Code: InvalidArgument</i> </p> </li> <li> <p>
// <i>Cause: Invalid argument.</i> </p> </li> </ul> </li> <li> <ul> <li> <p>
// <i>HTTP Error: HTTP 400 Bad Request</i> </p> </li> <li> <p> <i>Code:
// TooManyConfigurations</i> </p> </li> <li> <p> <i>Cause: You are attempting to
// create a new configuration but have already reached the 1,000-configuration
// limit.</i> </p> </li> </ul> </li> <li> <ul> <li> <p> <i>HTTP Error: HTTP 403
// Forbidden</i> </p> </li> <li> <p> <i>Code: AccessDenied</i> </p> </li> <li> <p>
// <i>Cause: You are not the owner of the specified bucket, or you do not have the
// s3:PutAnalyticsConfiguration bucket permission to set the configuration on the
// bucket.</i> </p> </li> </ul> </li> </ul> <p class="title"> <b>Related
// Resources</b> </p> <ul> <li> <p> </p> </li> <li> <p> </p> </li> <li> <p> </p>
// </li> </ul>
func (c *Client) PutBucketAnalyticsConfiguration(ctx context.Context, params *PutBucketAnalyticsConfigurationInput, optFns ...func(*Options)) (*PutBucketAnalyticsConfigurationOutput, error) {
	stack := middleware.NewStack("PutBucketAnalyticsConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpPutBucketAnalyticsConfigurationMiddlewares(stack)
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
	addOpPutBucketAnalyticsConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketAnalyticsConfiguration(options.Region), middleware.Before)
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
			OperationName: "PutBucketAnalyticsConfiguration",
			Err:           err,
		}
	}
	out := result.(*PutBucketAnalyticsConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketAnalyticsConfigurationInput struct {
	// The ID that identifies the analytics configuration.
	Id *string
	// The configuration and any analyses for the analytics filter.
	AnalyticsConfiguration *types.AnalyticsConfiguration
	// The name of the bucket to which an analytics configuration is stored.
	Bucket *string
}

type PutBucketAnalyticsConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpPutBucketAnalyticsConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpPutBucketAnalyticsConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketAnalyticsConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketAnalyticsConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketAnalyticsConfiguration",
	}
}
