// Code generated by "go generate". Please don't change this file directly.

package ses

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// MailManagerTrafficPolicy_IngressIpToEvaluate AWS CloudFormation Resource (AWS::SES::MailManagerTrafficPolicy.IngressIpToEvaluate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressiptoevaluate.html
type MailManagerTrafficPolicy_IngressIpToEvaluate struct {

	// Attribute AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagertrafficpolicy-ingressiptoevaluate.html#cfn-ses-mailmanagertrafficpolicy-ingressiptoevaluate-attribute
	Attribute string `json:"Attribute"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *MailManagerTrafficPolicy_IngressIpToEvaluate) AWSCloudFormationType() string {
	return "AWS::SES::MailManagerTrafficPolicy.IngressIpToEvaluate"
}
