// Code generated by "go generate". Please don't change this file directly.

package bedrock

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// PromptVersion_PromptTemplateConfiguration AWS CloudFormation Resource (AWS::Bedrock::PromptVersion.PromptTemplateConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-prompttemplateconfiguration.html
type PromptVersion_PromptTemplateConfiguration struct {

	// Chat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-prompttemplateconfiguration.html#cfn-bedrock-promptversion-prompttemplateconfiguration-chat
	Chat *PromptVersion_ChatPromptTemplateConfiguration `json:"Chat,omitempty"`

	// Text AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-prompttemplateconfiguration.html#cfn-bedrock-promptversion-prompttemplateconfiguration-text
	Text *PromptVersion_TextPromptTemplateConfiguration `json:"Text,omitempty"`

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
func (r *PromptVersion_PromptTemplateConfiguration) AWSCloudFormationType() string {
	return "AWS::Bedrock::PromptVersion.PromptTemplateConfiguration"
}
