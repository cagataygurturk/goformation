// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_TotalOptions AWS CloudFormation Resource (AWS::QuickSight::Template.TotalOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totaloptions.html
type Template_TotalOptions struct {

	// CustomLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totaloptions.html#cfn-quicksight-template-totaloptions-customlabel
	CustomLabel *string `json:"CustomLabel,omitempty"`

	// Placement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totaloptions.html#cfn-quicksight-template-totaloptions-placement
	Placement *string `json:"Placement,omitempty"`

	// ScrollStatus AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totaloptions.html#cfn-quicksight-template-totaloptions-scrollstatus
	ScrollStatus *string `json:"ScrollStatus,omitempty"`

	// TotalAggregationOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totaloptions.html#cfn-quicksight-template-totaloptions-totalaggregationoptions
	TotalAggregationOptions []Template_TotalAggregationOption `json:"TotalAggregationOptions,omitempty"`

	// TotalCellStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totaloptions.html#cfn-quicksight-template-totaloptions-totalcellstyle
	TotalCellStyle *Template_TableCellStyle `json:"TotalCellStyle,omitempty"`

	// TotalsVisibility AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totaloptions.html#cfn-quicksight-template-totaloptions-totalsvisibility
	TotalsVisibility interface{} `json:"TotalsVisibility,omitempty"`

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
func (r *Template_TotalOptions) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.TotalOptions"
}
