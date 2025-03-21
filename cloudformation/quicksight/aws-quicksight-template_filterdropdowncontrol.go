// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_FilterDropDownControl AWS CloudFormation Resource (AWS::QuickSight::Template.FilterDropDownControl)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html
type Template_FilterDropDownControl struct {

	// CascadingControlConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-cascadingcontrolconfiguration
	CascadingControlConfiguration *Template_CascadingControlConfiguration `json:"CascadingControlConfiguration,omitempty"`

	// CommitMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-commitmode
	CommitMode *string `json:"CommitMode,omitempty"`

	// DisplayOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-displayoptions
	DisplayOptions *Template_DropDownControlDisplayOptions `json:"DisplayOptions,omitempty"`

	// FilterControlId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-filtercontrolid
	FilterControlId string `json:"FilterControlId"`

	// SelectableValues AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-selectablevalues
	SelectableValues *Template_FilterSelectableValues `json:"SelectableValues,omitempty"`

	// SourceFilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-sourcefilterid
	SourceFilterId string `json:"SourceFilterId"`

	// Title AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-title
	Title string `json:"Title"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-filterdropdowncontrol.html#cfn-quicksight-template-filterdropdowncontrol-type
	Type *string `json:"Type,omitempty"`

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
func (r *Template_FilterDropDownControl) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.FilterDropDownControl"
}
