// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fms

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

var _ aws.Config
var _ = awsutil.Prettify

// Details of the resource that is not protected by the policy.
type ComplianceViolator struct {
	_ struct{} `type:"structure"`

	// The resource ID.
	ResourceId *string `min:"1" type:"string"`

	// The resource type. This is in the format shown in the AWS Resource Types
	// Reference (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html).
	// For example: AWS::ElasticLoadBalancingV2::LoadBalancer or AWS::CloudFront::Distribution.
	ResourceType *string `min:"1" type:"string"`

	// The reason that the resource is not protected by the policy.
	ViolationReason ViolationReason `type:"string" enum:"true"`
}

// String returns the string representation
func (s ComplianceViolator) String() string {
	return awsutil.Prettify(s)
}

// Describes the compliance status for the account. An account is considered
// noncompliant if it includes resources that are not protected by the specified
// policy or that don't comply with the policy.
type EvaluationResult struct {
	_ struct{} `type:"structure"`

	// Describes an AWS account's compliance with the AWS Firewall Manager policy.
	ComplianceStatus PolicyComplianceStatusType `type:"string" enum:"true"`

	// Indicates that over 100 resources are noncompliant with the AWS Firewall
	// Manager policy.
	EvaluationLimitExceeded *bool `type:"boolean"`

	// The number of resources that are noncompliant with the specified policy.
	// For AWS WAF and Shield Advanced policies, a resource is considered noncompliant
	// if it is not associated with the policy. For security group policies, a resource
	// is considered noncompliant if it doesn't comply with the rules of the policy
	// and remediation is disabled or not possible.
	ViolatorCount *int64 `type:"long"`
}

// String returns the string representation
func (s EvaluationResult) String() string {
	return awsutil.Prettify(s)
}

// An AWS Firewall Manager policy.
type Policy struct {
	_ struct{} `type:"structure"`

	// Specifies the AWS account IDs and AWS Organizations organizational units
	// (OUs) to exclude from the policy. Specifying an OU is the equivalent of specifying
	// all accounts in the OU and in any of its child OUs, including any child OUs
	// and accounts that are added at a later time.
	//
	// You can specify inclusions or exclusions, but not both. If you specify an
	// IncludeMap, AWS Firewall Manager applies the policy to all accounts specified
	// by the IncludeMap, and does not evaluate any ExcludeMap specifications. If
	// you do not specify an IncludeMap, then Firewall Manager applies the policy
	// to all accounts except for those specified by the ExcludeMap.
	//
	// You can specify account IDs, OUs, or a combination:
	//
	//    * Specify account IDs by setting the key to ACCOUNT. For example, the
	//    following is a valid map: {“ACCOUNT” : [“accountID1”, “accountID2”]}.
	//
	//    * Specify OUs by setting the key to ORG_UNIT. For example, the following
	//    is a valid map: {“ORG_UNIT” : [“ouid111”, “ouid112”]}.
	//
	//    * Specify accounts and OUs together in a single map, separated with a
	//    comma. For example, the following is a valid map: {“ACCOUNT” : [“accountID1”,
	//    “accountID2”], “ORG_UNIT” : [“ouid111”, “ouid112”]}.
	ExcludeMap map[string][]string `type:"map"`

	// If set to True, resources with the tags that are specified in the ResourceTag
	// array are not in scope of the policy. If set to False, and the ResourceTag
	// array is not null, only resources with the specified tags are in scope of
	// the policy.
	//
	// ExcludeResourceTags is a required field
	ExcludeResourceTags *bool `type:"boolean" required:"true"`

	// Specifies the AWS account IDs and AWS Organizations organizational units
	// (OUs) to include in the policy. Specifying an OU is the equivalent of specifying
	// all accounts in the OU and in any of its child OUs, including any child OUs
	// and accounts that are added at a later time.
	//
	// You can specify inclusions or exclusions, but not both. If you specify an
	// IncludeMap, AWS Firewall Manager applies the policy to all accounts specified
	// by the IncludeMap, and does not evaluate any ExcludeMap specifications. If
	// you do not specify an IncludeMap, then Firewall Manager applies the policy
	// to all accounts except for those specified by the ExcludeMap.
	//
	// You can specify account IDs, OUs, or a combination:
	//
	//    * Specify account IDs by setting the key to ACCOUNT. For example, the
	//    following is a valid map: {“ACCOUNT” : [“accountID1”, “accountID2”]}.
	//
	//    * Specify OUs by setting the key to ORG_UNIT. For example, the following
	//    is a valid map: {“ORG_UNIT” : [“ouid111”, “ouid112”]}.
	//
	//    * Specify accounts and OUs together in a single map, separated with a
	//    comma. For example, the following is a valid map: {“ACCOUNT” : [“accountID1”,
	//    “accountID2”], “ORG_UNIT” : [“ouid111”, “ouid112”]}.
	IncludeMap map[string][]string `type:"map"`

	// The ID of the AWS Firewall Manager policy.
	PolicyId *string `min:"36" type:"string"`

	// The friendly name of the AWS Firewall Manager policy.
	//
	// PolicyName is a required field
	PolicyName *string `min:"1" type:"string" required:"true"`

	// A unique identifier for each update to the policy. When issuing a PutPolicy
	// request, the PolicyUpdateToken in the request must match the PolicyUpdateToken
	// of the current policy version. To get the PolicyUpdateToken of the current
	// policy version, use a GetPolicy request.
	PolicyUpdateToken *string `min:"1" type:"string"`

	// Indicates if the policy should be automatically applied to new resources.
	//
	// RemediationEnabled is a required field
	RemediationEnabled *bool `type:"boolean" required:"true"`

	// An array of ResourceTag objects.
	ResourceTags []ResourceTag `type:"list"`

	// The type of resource protected by or in scope of the policy. This is in the
	// format shown in the AWS Resource Types Reference (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html).
	// For AWS WAF and Shield Advanced, examples include AWS::ElasticLoadBalancingV2::LoadBalancer
	// and AWS::CloudFront::Distribution. For a security group common policy, valid
	// values are AWS::EC2::NetworkInterface and AWS::EC2::Instance. For a security
	// group content audit policy, valid values are AWS::EC2::SecurityGroup, AWS::EC2::NetworkInterface,
	// and AWS::EC2::Instance. For a security group usage audit policy, the value
	// is AWS::EC2::SecurityGroup.
	//
	// ResourceType is a required field
	ResourceType *string `min:"1" type:"string" required:"true"`

	// An array of ResourceType.
	ResourceTypeList []string `type:"list"`

	// Details about the security service that is being used to protect the resources.
	//
	// SecurityServicePolicyData is a required field
	SecurityServicePolicyData *SecurityServicePolicyData `type:"structure" required:"true"`
}

// String returns the string representation
func (s Policy) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Policy) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Policy"}

	if s.ExcludeResourceTags == nil {
		invalidParams.Add(aws.NewErrParamRequired("ExcludeResourceTags"))
	}
	if s.PolicyId != nil && len(*s.PolicyId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyId", 36))
	}

	if s.PolicyName == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyName", 1))
	}
	if s.PolicyUpdateToken != nil && len(*s.PolicyUpdateToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyUpdateToken", 1))
	}

	if s.RemediationEnabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("RemediationEnabled"))
	}

	if s.ResourceType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceType"))
	}
	if s.ResourceType != nil && len(*s.ResourceType) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceType", 1))
	}

	if s.SecurityServicePolicyData == nil {
		invalidParams.Add(aws.NewErrParamRequired("SecurityServicePolicyData"))
	}
	if s.ResourceTags != nil {
		for i, v := range s.ResourceTags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ResourceTags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SecurityServicePolicyData != nil {
		if err := s.SecurityServicePolicyData.Validate(); err != nil {
			invalidParams.AddNested("SecurityServicePolicyData", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Describes the noncompliant resources in a member account for a specific AWS
// Firewall Manager policy. A maximum of 100 entries are displayed. If more
// than 100 resources are noncompliant, EvaluationLimitExceeded is set to True.
type PolicyComplianceDetail struct {
	_ struct{} `type:"structure"`

	// Indicates if over 100 resources are noncompliant with the AWS Firewall Manager
	// policy.
	EvaluationLimitExceeded *bool `type:"boolean"`

	// A timestamp that indicates when the returned information should be considered
	// out of date.
	ExpiredAt *time.Time `type:"timestamp"`

	// Details about problems with dependent services, such as AWS WAF or AWS Config,
	// that are causing a resource to be noncompliant. The details include the name
	// of the dependent service and the error message received that indicates the
	// problem with the service.
	IssueInfoMap map[string]string `type:"map"`

	// The AWS account ID.
	MemberAccount *string `min:"1" type:"string"`

	// The ID of the AWS Firewall Manager policy.
	PolicyId *string `min:"36" type:"string"`

	// The AWS account that created the AWS Firewall Manager policy.
	PolicyOwner *string `min:"1" type:"string"`

	// An array of resources that aren't protected by the AWS WAF or Shield Advanced
	// policy or that aren't in compliance with the security group policy.
	Violators []ComplianceViolator `type:"list"`
}

// String returns the string representation
func (s PolicyComplianceDetail) String() string {
	return awsutil.Prettify(s)
}

// Indicates whether the account is compliant with the specified policy. An
// account is considered noncompliant if it includes resources that are not
// protected by the policy, for AWS WAF and Shield Advanced policies, or that
// are noncompliant with the policy, for security group policies.
type PolicyComplianceStatus struct {
	_ struct{} `type:"structure"`

	// An array of EvaluationResult objects.
	EvaluationResults []EvaluationResult `type:"list"`

	// Details about problems with dependent services, such as AWS WAF or AWS Config,
	// that are causing a resource to be noncompliant. The details include the name
	// of the dependent service and the error message received that indicates the
	// problem with the service.
	IssueInfoMap map[string]string `type:"map"`

	// Timestamp of the last update to the EvaluationResult objects.
	LastUpdated *time.Time `type:"timestamp"`

	// The member account ID.
	MemberAccount *string `min:"1" type:"string"`

	// The ID of the AWS Firewall Manager policy.
	PolicyId *string `min:"36" type:"string"`

	// The friendly name of the AWS Firewall Manager policy.
	PolicyName *string `min:"1" type:"string"`

	// The AWS account that created the AWS Firewall Manager policy.
	PolicyOwner *string `min:"1" type:"string"`
}

// String returns the string representation
func (s PolicyComplianceStatus) String() string {
	return awsutil.Prettify(s)
}

// Details of the AWS Firewall Manager policy.
type PolicySummary struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the specified policy.
	PolicyArn *string `min:"1" type:"string"`

	// The ID of the specified policy.
	PolicyId *string `min:"36" type:"string"`

	// The friendly name of the specified policy.
	PolicyName *string `min:"1" type:"string"`

	// Indicates if the policy should be automatically applied to new resources.
	RemediationEnabled *bool `type:"boolean"`

	// The type of resource protected by or in scope of the policy. This is in the
	// format shown in the AWS Resource Types Reference (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html).
	// For AWS WAF and Shield Advanced, examples include AWS::ElasticLoadBalancingV2::LoadBalancer
	// and AWS::CloudFront::Distribution. For a security group common policy, valid
	// values are AWS::EC2::NetworkInterface and AWS::EC2::Instance. For a security
	// group content audit policy, valid values are AWS::EC2::SecurityGroup, AWS::EC2::NetworkInterface,
	// and AWS::EC2::Instance. For a security group usage audit policy, the value
	// is AWS::EC2::SecurityGroup.
	ResourceType *string `min:"1" type:"string"`

	// The service that the policy is using to protect the resources. This specifies
	// the type of policy that is created, either an AWS WAF policy, a Shield Advanced
	// policy, or a security group policy.
	SecurityServiceType SecurityServiceType `type:"string" enum:"true"`
}

// String returns the string representation
func (s PolicySummary) String() string {
	return awsutil.Prettify(s)
}

// The resource tags that AWS Firewall Manager uses to determine if a particular
// resource should be included or excluded from the AWS Firewall Manager policy.
// Tags enable you to categorize your AWS resources in different ways, for example,
// by purpose, owner, or environment. Each tag consists of a key and an optional
// value. Firewall Manager combines the tags with "AND" so that, if you add
// more than one tag to a policy scope, a resource must have all the specified
// tags to be included or excluded. For more information, see Working with Tag
// Editor (https://docs.aws.amazon.com/awsconsolehelpdocs/latest/gsg/tag-editor.html).
type ResourceTag struct {
	_ struct{} `type:"structure"`

	// The resource tag key.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// The resource tag value.
	Value *string `type:"string"`
}

// String returns the string representation
func (s ResourceTag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResourceTag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResourceTag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Details about the security service that is being used to protect the resources.
type SecurityServicePolicyData struct {
	_ struct{} `type:"structure"`

	// Details about the service that are specific to the service type, in JSON
	// format. For service type SHIELD_ADVANCED, this is an empty string.
	//
	//    * Example: WAFV2 "ManagedServiceData": "{\"type\":\"WAFV2\",\"defaultAction\":{\"type\":\"ALLOW\"},\"preProcessRuleGroups\":[{\"managedRuleGroupIdentifier\":null,\"ruleGroupArn\":\"rulegrouparn\",\"overrideAction\":{\"type\":\"COUNT\"},\"excludedRules\":[{\"name\":\"EntityName\"}],\"ruleGroupType\":\"RuleGroup\"}],\"postProcessRuleGroups\":[{\"managedRuleGroupIdentifier\":{\"managedRuleGroupName\":\"AWSManagedRulesAdminProtectionRuleSet\",\"vendor\":\"AWS\"},\"ruleGroupArn\":\"rulegrouparn\",\"overrideAction\":{\"type\":\"NONE\"},\"excludedRules\":[],\"ruleGroupType\":\"ManagedRuleGroup\"}],\"overrideCustomerWebACLAssociation\":false}"
	//
	//    * Example: WAF Classic "ManagedServiceData": "{\"type\": \"WAF\", \"ruleGroups\":
	//    [{\"id\": \"12345678-1bcd-9012-efga-0987654321ab\", \"overrideAction\"
	//    : {\"type\": \"COUNT\"}}], \"defaultAction\": {\"type\": \"BLOCK\"}}
	//
	//    * Example: SECURITY_GROUPS_COMMON "SecurityServicePolicyData":{"Type":"SECURITY_GROUPS_COMMON","ManagedServiceData":"{\"type\":\"SECURITY_GROUPS_COMMON\",\"revertManualSecurityGroupChanges\":false,\"exclusiveResourceSecurityGroupManagement\":false,
	//    \"applyToAllEC2InstanceENIs\":false,\"securityGroups\":[{\"id\":\" sg-000e55995d61a06bd\"}]}"},"RemediationEnabled":false,"ResourceType":"AWS::EC2::NetworkInterface"}
	//
	//    * Example: SECURITY_GROUPS_CONTENT_AUDIT "SecurityServicePolicyData":{"Type":"SECURITY_GROUPS_CONTENT_AUDIT","ManagedServiceData":"{\"type\":\"SECURITY_GROUPS_CONTENT_AUDIT\",\"securityGroups\":[{\"id\":\"
	//    sg-000e55995d61a06bd \"}],\"securityGroupAction\":{\"type\":\"ALLOW\"}}"},"RemediationEnabled":false,"ResourceType":"AWS::EC2::NetworkInterface"}
	//    The security group action for content audit can be ALLOW or DENY. For
	//    ALLOW, all in-scope security group rules must be within the allowed range
	//    of the policy's security group rules. For DENY, all in-scope security
	//    group rules must not contain a value or a range that matches a rule value
	//    or range in the policy security group.
	//
	//    * Example: SECURITY_GROUPS_USAGE_AUDIT "SecurityServicePolicyData":{"Type":"SECURITY_GROUPS_USAGE_AUDIT","ManagedServiceData":"{\"type\":\"SECURITY_GROUPS_USAGE_AUDIT\",\"deleteUnusedSecurityGroups\":true,\"coalesceRedundantSecurityGroups\":true}"},"RemediationEnabled":false,"Resou
	//    rceType":"AWS::EC2::SecurityGroup"}
	ManagedServiceData *string `min:"1" type:"string"`

	// The service that the policy is using to protect the resources. This specifies
	// the type of policy that is created, either an AWS WAF policy, a Shield Advanced
	// policy, or a security group policy. For security group policies, Firewall
	// Manager supports one security group for each common policy and for each content
	// audit policy. This is an adjustable limit that you can increase by contacting
	// AWS Support.
	//
	// Type is a required field
	Type SecurityServiceType `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s SecurityServicePolicyData) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SecurityServicePolicyData) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SecurityServicePolicyData"}
	if s.ManagedServiceData != nil && len(*s.ManagedServiceData) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ManagedServiceData", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A collection of key:value pairs associated with an AWS resource. The key:value
// pair can be anything you define. Typically, the tag key represents a category
// (such as "environment") and the tag value represents a specific value within
// that category (such as "test," "development," or "production"). You can add
// up to 50 tags to each AWS resource.
type Tag struct {
	_ struct{} `type:"structure"`

	// Part of the key:value pair that defines a tag. You can use a tag key to describe
	// a category of information, such as "customer." Tag keys are case-sensitive.
	//
	// Key is a required field
	Key *string `min:"1" type:"string" required:"true"`

	// Part of the key:value pair that defines a tag. You can use a tag value to
	// describe a specific value within a category, such as "companyA" or "companyB."
	// Tag values are case-sensitive.
	//
	// Value is a required field
	Value *string `type:"string" required:"true"`
}

// String returns the string representation
func (s Tag) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *Tag) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "Tag"}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
