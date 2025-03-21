// Code generated by "go generate". Please don't change this file directly.

package logs

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Transformer_ParseVPC AWS CloudFormation Resource (AWS::Logs::Transformer.ParseVPC)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsevpc.html
type Transformer_ParseVPC struct {

	// Source AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-transformer-parsevpc.html#cfn-logs-transformer-parsevpc-source
	Source *string `json:"Source,omitempty"`

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
func (r *Transformer_ParseVPC) AWSCloudFormationType() string {
	return "AWS::Logs::Transformer.ParseVPC"
}
