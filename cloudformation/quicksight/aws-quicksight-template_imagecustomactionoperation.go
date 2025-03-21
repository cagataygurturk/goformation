// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_ImageCustomActionOperation AWS CloudFormation Resource (AWS::QuickSight::Template.ImageCustomActionOperation)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagecustomactionoperation.html
type Template_ImageCustomActionOperation struct {

	// NavigationOperation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagecustomactionoperation.html#cfn-quicksight-template-imagecustomactionoperation-navigationoperation
	NavigationOperation *Template_CustomActionNavigationOperation `json:"NavigationOperation,omitempty"`

	// SetParametersOperation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagecustomactionoperation.html#cfn-quicksight-template-imagecustomactionoperation-setparametersoperation
	SetParametersOperation *Template_CustomActionSetParametersOperation `json:"SetParametersOperation,omitempty"`

	// URLOperation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-imagecustomactionoperation.html#cfn-quicksight-template-imagecustomactionoperation-urloperation
	URLOperation *Template_CustomActionURLOperation `json:"URLOperation,omitempty"`

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
func (r *Template_ImageCustomActionOperation) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.ImageCustomActionOperation"
}
