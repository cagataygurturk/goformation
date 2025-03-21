// Code generated by "go generate". Please don't change this file directly.

package resiliencehub

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// ResiliencyPolicy_PolicyMap AWS CloudFormation Resource (AWS::ResilienceHub::ResiliencyPolicy.PolicyMap)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-resiliencypolicy-policymap.html
type ResiliencyPolicy_PolicyMap struct {

	// AZ AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-resiliencypolicy-policymap.html#cfn-resiliencehub-resiliencypolicy-policymap-az
	AZ *ResiliencyPolicy_FailurePolicy `json:"AZ"`

	// Hardware AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-resiliencypolicy-policymap.html#cfn-resiliencehub-resiliencypolicy-policymap-hardware
	Hardware *ResiliencyPolicy_FailurePolicy `json:"Hardware"`

	// Region AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-resiliencypolicy-policymap.html#cfn-resiliencehub-resiliencypolicy-policymap-region
	Region *ResiliencyPolicy_FailurePolicy `json:"Region,omitempty"`

	// Software AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resiliencehub-resiliencypolicy-policymap.html#cfn-resiliencehub-resiliencypolicy-policymap-software
	Software *ResiliencyPolicy_FailurePolicy `json:"Software"`

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
func (r *ResiliencyPolicy_PolicyMap) AWSCloudFormationType() string {
	return "AWS::ResilienceHub::ResiliencyPolicy.PolicyMap"
}
