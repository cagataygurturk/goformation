// Code generated by "go generate". Please don't change this file directly.

package bedrock

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// FlowVersion_GuardrailConfiguration AWS CloudFormation Resource (AWS::Bedrock::FlowVersion.GuardrailConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-guardrailconfiguration.html
type FlowVersion_GuardrailConfiguration struct {

	// GuardrailIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-guardrailconfiguration.html#cfn-bedrock-flowversion-guardrailconfiguration-guardrailidentifier
	GuardrailIdentifier *string `json:"GuardrailIdentifier,omitempty"`

	// GuardrailVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-guardrailconfiguration.html#cfn-bedrock-flowversion-guardrailconfiguration-guardrailversion
	GuardrailVersion *string `json:"GuardrailVersion,omitempty"`

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
func (r *FlowVersion_GuardrailConfiguration) AWSCloudFormationType() string {
	return "AWS::Bedrock::FlowVersion.GuardrailConfiguration"
}
