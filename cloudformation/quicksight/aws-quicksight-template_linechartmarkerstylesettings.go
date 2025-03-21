// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_LineChartMarkerStyleSettings AWS CloudFormation Resource (AWS::QuickSight::Template.LineChartMarkerStyleSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartmarkerstylesettings.html
type Template_LineChartMarkerStyleSettings struct {

	// MarkerColor AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartmarkerstylesettings.html#cfn-quicksight-template-linechartmarkerstylesettings-markercolor
	MarkerColor *string `json:"MarkerColor,omitempty"`

	// MarkerShape AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartmarkerstylesettings.html#cfn-quicksight-template-linechartmarkerstylesettings-markershape
	MarkerShape *string `json:"MarkerShape,omitempty"`

	// MarkerSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartmarkerstylesettings.html#cfn-quicksight-template-linechartmarkerstylesettings-markersize
	MarkerSize *string `json:"MarkerSize,omitempty"`

	// MarkerVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-linechartmarkerstylesettings.html#cfn-quicksight-template-linechartmarkerstylesettings-markervisibility
	MarkerVisibility interface{} `json:"MarkerVisibility,omitempty"`

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
func (r *Template_LineChartMarkerStyleSettings) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.LineChartMarkerStyleSettings"
}
