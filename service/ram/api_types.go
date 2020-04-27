// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ram

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Describes a principal for use with AWS Resource Access Manager.
type Principal struct {
	_ struct{} `type:"structure"`

	// The time when the principal was associated with the resource share.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp"`

	// Indicates whether the principal belongs to the same AWS organization as the
	// AWS account that owns the resource share.
	External *bool `locationName:"external" type:"boolean"`

	// The ID of the principal.
	Id *string `locationName:"id" type:"string"`

	// The time when the association was last updated.
	LastUpdatedTime *time.Time `locationName:"lastUpdatedTime" type:"timestamp"`

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string `locationName:"resourceShareArn" type:"string"`
}

// String returns the string representation
func (s Principal) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Principal) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.External != nil {
		v := *s.External

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "external", protocol.BoolValue(v), metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.LastUpdatedTime != nil {
		v := *s.LastUpdatedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.ResourceShareArn != nil {
		v := *s.ResourceShareArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Describes a resource associated with a resource share.
type Resource struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the resource.
	Arn *string `locationName:"arn" type:"string"`

	// The time when the resource was associated with the resource share.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp"`

	// The time when the association was last updated.
	LastUpdatedTime *time.Time `locationName:"lastUpdatedTime" type:"timestamp"`

	// The ARN of the resource group. This value is returned only if the resource
	// is a resource group.
	ResourceGroupArn *string `locationName:"resourceGroupArn" type:"string"`

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string `locationName:"resourceShareArn" type:"string"`

	// The status of the resource.
	Status ResourceStatus `locationName:"status" type:"string" enum:"true"`

	// A message about the status of the resource.
	StatusMessage *string `locationName:"statusMessage" type:"string"`

	// The resource type.
	Type *string `locationName:"type" type:"string"`
}

// String returns the string representation
func (s Resource) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Resource) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.LastUpdatedTime != nil {
		v := *s.LastUpdatedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.ResourceGroupArn != nil {
		v := *s.ResourceGroupArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceGroupArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareArn != nil {
		v := *s.ResourceShareArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusMessage != nil {
		v := *s.StatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Type != nil {
		v := *s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Describes a resource share.
type ResourceShare struct {
	_ struct{} `type:"structure"`

	// Indicates whether principals outside your AWS organization can be associated
	// with a resource share.
	AllowExternalPrincipals *bool `locationName:"allowExternalPrincipals" type:"boolean"`

	// The time when the resource share was created.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp"`

	// Indicates how the resource share was created. Possible values include:
	//
	//    * CREATED_FROM_POLICY - Indicates that the resource share was created
	//    from an AWS Identity and Access Management (AWS IAM) policy attached to
	//    a resource. These resource shares are visible only to the AWS account
	//    that created it. They cannot be modified in AWS RAM.
	//
	//    * PROMOTING_TO_STANDARD - The resource share is in the process of being
	//    promoted. For more information, see PromoteResourceShareCreatedFromPolicy.
	//
	//    * STANDARD - Indicates that the resource share was created in AWS RAM
	//    using the console or APIs. These resource shares are visible to all principals.
	//    They can be modified in AWS RAM.
	FeatureSet ResourceShareFeatureSet `locationName:"featureSet" type:"string" enum:"true"`

	// The time when the resource share was last updated.
	LastUpdatedTime *time.Time `locationName:"lastUpdatedTime" type:"timestamp"`

	// The name of the resource share.
	Name *string `locationName:"name" type:"string"`

	// The ID of the AWS account that owns the resource share.
	OwningAccountId *string `locationName:"owningAccountId" type:"string"`

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string `locationName:"resourceShareArn" type:"string"`

	// The status of the resource share.
	Status ResourceShareStatus `locationName:"status" type:"string" enum:"true"`

	// A message about the status of the resource share.
	StatusMessage *string `locationName:"statusMessage" type:"string"`

	// The tags for the resource share.
	Tags []Tag `locationName:"tags" type:"list"`
}

// String returns the string representation
func (s ResourceShare) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ResourceShare) MarshalFields(e protocol.FieldEncoder) error {
	if s.AllowExternalPrincipals != nil {
		v := *s.AllowExternalPrincipals

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "allowExternalPrincipals", protocol.BoolValue(v), metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if len(s.FeatureSet) > 0 {
		v := s.FeatureSet

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "featureSet", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.LastUpdatedTime != nil {
		v := *s.LastUpdatedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OwningAccountId != nil {
		v := *s.OwningAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "owningAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareArn != nil {
		v := *s.ResourceShareArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusMessage != nil {
		v := *s.StatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

// Describes an association with a resource share.
type ResourceShareAssociation struct {
	_ struct{} `type:"structure"`

	// The associated entity. For resource associations, this is the ARN of the
	// resource. For principal associations, this is the ID of an AWS account or
	// the ARN of an OU or organization from AWS Organizations.
	AssociatedEntity *string `locationName:"associatedEntity" type:"string"`

	// The association type.
	AssociationType ResourceShareAssociationType `locationName:"associationType" type:"string" enum:"true"`

	// The time when the association was created.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp"`

	// Indicates whether the principal belongs to the same AWS organization as the
	// AWS account that owns the resource share.
	External *bool `locationName:"external" type:"boolean"`

	// The time when the association was last updated.
	LastUpdatedTime *time.Time `locationName:"lastUpdatedTime" type:"timestamp"`

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string `locationName:"resourceShareArn" type:"string"`

	// The name of the resource share.
	ResourceShareName *string `locationName:"resourceShareName" type:"string"`

	// The status of the association.
	Status ResourceShareAssociationStatus `locationName:"status" type:"string" enum:"true"`

	// A message about the status of the association.
	StatusMessage *string `locationName:"statusMessage" type:"string"`
}

// String returns the string representation
func (s ResourceShareAssociation) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ResourceShareAssociation) MarshalFields(e protocol.FieldEncoder) error {
	if s.AssociatedEntity != nil {
		v := *s.AssociatedEntity

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "associatedEntity", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.AssociationType) > 0 {
		v := s.AssociationType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "associationType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.External != nil {
		v := *s.External

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "external", protocol.BoolValue(v), metadata)
	}
	if s.LastUpdatedTime != nil {
		v := *s.LastUpdatedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.ResourceShareArn != nil {
		v := *s.ResourceShareArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareName != nil {
		v := *s.ResourceShareName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.StatusMessage != nil {
		v := *s.StatusMessage

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "statusMessage", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Describes an invitation to join a resource share.
type ResourceShareInvitation struct {
	_ struct{} `type:"structure"`

	// The date and time when the invitation was sent.
	InvitationTimestamp *time.Time `locationName:"invitationTimestamp" type:"timestamp"`

	// The ID of the AWS account that received the invitation.
	ReceiverAccountId *string `locationName:"receiverAccountId" type:"string"`

	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string `locationName:"resourceShareArn" type:"string"`

	// To view the resources associated with a pending resource share invitation,
	// use ListPendingInvitationResources (https://docs.aws.amazon.com/ram/latest/APIReference/API_ListPendingInvitationResources.html).
	ResourceShareAssociations []ResourceShareAssociation `locationName:"resourceShareAssociations" deprecated:"true" type:"list"`

	// The Amazon Resource Name (ARN) of the invitation.
	ResourceShareInvitationArn *string `locationName:"resourceShareInvitationArn" type:"string"`

	// The name of the resource share.
	ResourceShareName *string `locationName:"resourceShareName" type:"string"`

	// The ID of the AWS account that sent the invitation.
	SenderAccountId *string `locationName:"senderAccountId" type:"string"`

	// The status of the invitation.
	Status ResourceShareInvitationStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ResourceShareInvitation) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ResourceShareInvitation) MarshalFields(e protocol.FieldEncoder) error {
	if s.InvitationTimestamp != nil {
		v := *s.InvitationTimestamp

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "invitationTimestamp",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.ReceiverAccountId != nil {
		v := *s.ReceiverAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "receiverAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareArn != nil {
		v := *s.ResourceShareArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareAssociations != nil {
		v := s.ResourceShareAssociations

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "resourceShareAssociations", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ResourceShareInvitationArn != nil {
		v := *s.ResourceShareInvitationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareInvitationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceShareName != nil {
		v := *s.ResourceShareName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceShareName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SenderAccountId != nil {
		v := *s.SenderAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "senderAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Status) > 0 {
		v := s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Information about an AWS RAM permission.
type ResourceSharePermissionDetail struct {
	_ struct{} `type:"structure"`

	// The ARN of the permission.
	Arn *string `locationName:"arn" type:"string"`

	// The date and time when the permission was created.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp"`

	// The identifier for the version of the permission that is set as the default
	// version.
	DefaultVersion *bool `locationName:"defaultVersion" type:"boolean"`

	// The date and time when the permission was last updated.
	LastUpdatedTime *time.Time `locationName:"lastUpdatedTime" type:"timestamp"`

	// The name of the permission.
	Name *string `locationName:"name" type:"string"`

	// The permission's effect and actions in JSON format. The effect indicates
	// whether the actions are allowed or denied. The actions list the API actions
	// to which the principal is granted or denied access.
	Permission *string `locationName:"permission" type:"string"`

	// The resource type to which the permission applies.
	ResourceType *string `locationName:"resourceType" type:"string"`

	// The identifier for the version of the permission.
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s ResourceSharePermissionDetail) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ResourceSharePermissionDetail) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.DefaultVersion != nil {
		v := *s.DefaultVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultVersion", protocol.BoolValue(v), metadata)
	}
	if s.LastUpdatedTime != nil {
		v := *s.LastUpdatedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Permission != nil {
		v := *s.Permission

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "permission", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceType != nil {
		v := *s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about a permission that is associated with a resource share.
type ResourceSharePermissionSummary struct {
	_ struct{} `type:"structure"`

	// The ARN of the permission.
	Arn *string `locationName:"arn" type:"string"`

	// The date and time when the permission was created.
	CreationTime *time.Time `locationName:"creationTime" type:"timestamp"`

	// The identifier for the version of the permission that is set as the default
	// version.
	DefaultVersion *bool `locationName:"defaultVersion" type:"boolean"`

	// The date and time when the permission was last updated.
	LastUpdatedTime *time.Time `locationName:"lastUpdatedTime" type:"timestamp"`

	// The name of the permission.
	Name *string `locationName:"name" type:"string"`

	// The type of resource to which the permission applies.
	ResourceType *string `locationName:"resourceType" type:"string"`

	// The current status of the permission.
	Status *string `locationName:"status" type:"string"`

	// The identifier for the version of the permission.
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s ResourceSharePermissionSummary) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ResourceSharePermissionSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CreationTime != nil {
		v := *s.CreationTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.DefaultVersion != nil {
		v := *s.DefaultVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "defaultVersion", protocol.BoolValue(v), metadata)
	}
	if s.LastUpdatedTime != nil {
		v := *s.LastUpdatedTime

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "lastUpdatedTime",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ResourceType != nil {
		v := *s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Status != nil {
		v := *s.Status

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "status", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about the shareable resource types and the AWS services to which
// they belong.
type ServiceNameAndResourceType struct {
	_ struct{} `type:"structure"`

	// The shareable resource types.
	ResourceType *string `locationName:"resourceType" type:"string"`

	// The name of the AWS services to which the resources belong.
	ServiceName *string `locationName:"serviceName" type:"string"`
}

// String returns the string representation
func (s ServiceNameAndResourceType) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ServiceNameAndResourceType) MarshalFields(e protocol.FieldEncoder) error {
	if s.ResourceType != nil {
		v := *s.ResourceType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "resourceType", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ServiceName != nil {
		v := *s.ServiceName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "serviceName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Information about a tag.
type Tag struct {
	_ struct{} `type:"structure"`

	// The key of the tag.
	Key *string `locationName:"key" type:"string"`

	// The value of the tag.
	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Tag) MarshalFields(e protocol.FieldEncoder) error {
	if s.Key != nil {
		v := *s.Key

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "key", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Used to filter information based on tags.
type TagFilter struct {
	_ struct{} `type:"structure"`

	// The tag key.
	TagKey *string `locationName:"tagKey" type:"string"`

	// The tag values.
	TagValues []string `locationName:"tagValues" type:"list"`
}

// String returns the string representation
func (s TagFilter) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s TagFilter) MarshalFields(e protocol.FieldEncoder) error {
	if s.TagKey != nil {
		v := *s.TagKey

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "tagKey", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TagValues != nil {
		v := s.TagValues

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "tagValues", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}
