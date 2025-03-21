// Code generated by "go generate". Please don't change this file directly.

package dynamodb

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// GlobalTable_WriteOnDemandThroughputSettings AWS CloudFormation Resource (AWS::DynamoDB::GlobalTable.WriteOnDemandThroughputSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-writeondemandthroughputsettings.html
type GlobalTable_WriteOnDemandThroughputSettings struct {

	// MaxWriteRequestUnits AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-writeondemandthroughputsettings.html#cfn-dynamodb-globaltable-writeondemandthroughputsettings-maxwriterequestunits
	MaxWriteRequestUnits *int `json:"MaxWriteRequestUnits,omitempty"`

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
func (r *GlobalTable_WriteOnDemandThroughputSettings) AWSCloudFormationType() string {
	return "AWS::DynamoDB::GlobalTable.WriteOnDemandThroughputSettings"
}
