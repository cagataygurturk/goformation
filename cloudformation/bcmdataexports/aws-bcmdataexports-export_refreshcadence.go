// Code generated by "go generate". Please don't change this file directly.

package bcmdataexports

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Export_RefreshCadence AWS CloudFormation Resource (AWS::BCMDataExports::Export.RefreshCadence)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-refreshcadence.html
type Export_RefreshCadence struct {

	// Frequency AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bcmdataexports-export-refreshcadence.html#cfn-bcmdataexports-export-refreshcadence-frequency
	Frequency string `json:"Frequency"`

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
func (r *Export_RefreshCadence) AWSCloudFormationType() string {
	return "AWS::BCMDataExports::Export.RefreshCadence"
}
