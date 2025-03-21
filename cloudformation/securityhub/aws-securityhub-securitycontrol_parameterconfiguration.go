// Code generated by "go generate". Please don't change this file directly.

package securityhub

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// SecurityControl_ParameterConfiguration AWS CloudFormation Resource (AWS::SecurityHub::SecurityControl.ParameterConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parameterconfiguration.html
type SecurityControl_ParameterConfiguration struct {

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parameterconfiguration.html#cfn-securityhub-securitycontrol-parameterconfiguration-value
	Value *SecurityControl_ParameterValue `json:"Value,omitempty"`

	// ValueType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securityhub-securitycontrol-parameterconfiguration.html#cfn-securityhub-securitycontrol-parameterconfiguration-valuetype
	ValueType string `json:"ValueType"`

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
func (r *SecurityControl_ParameterConfiguration) AWSCloudFormationType() string {
	return "AWS::SecurityHub::SecurityControl.ParameterConfiguration"
}
