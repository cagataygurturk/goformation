// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Template_NestedFilter AWS CloudFormation Resource (AWS::QuickSight::Template.NestedFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nestedfilter.html
type Template_NestedFilter struct {

	// Column AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nestedfilter.html#cfn-quicksight-template-nestedfilter-column
	Column *Template_ColumnIdentifier `json:"Column"`

	// FilterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nestedfilter.html#cfn-quicksight-template-nestedfilter-filterid
	FilterId string `json:"FilterId"`

	// IncludeInnerSet AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nestedfilter.html#cfn-quicksight-template-nestedfilter-includeinnerset
	IncludeInnerSet bool `json:"IncludeInnerSet"`

	// InnerFilter AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-nestedfilter.html#cfn-quicksight-template-nestedfilter-innerfilter
	InnerFilter *Template_InnerFilter `json:"InnerFilter"`

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
func (r *Template_NestedFilter) AWSCloudFormationType() string {
	return "AWS::QuickSight::Template.NestedFilter"
}
