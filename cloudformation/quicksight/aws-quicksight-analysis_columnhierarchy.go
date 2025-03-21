// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_ColumnHierarchy AWS CloudFormation Resource (AWS::QuickSight::Analysis.ColumnHierarchy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnhierarchy.html
type Analysis_ColumnHierarchy struct {

	// DateTimeHierarchy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnhierarchy.html#cfn-quicksight-analysis-columnhierarchy-datetimehierarchy
	DateTimeHierarchy *Analysis_DateTimeHierarchy `json:"DateTimeHierarchy,omitempty"`

	// ExplicitHierarchy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnhierarchy.html#cfn-quicksight-analysis-columnhierarchy-explicithierarchy
	ExplicitHierarchy *Analysis_ExplicitHierarchy `json:"ExplicitHierarchy,omitempty"`

	// PredefinedHierarchy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-columnhierarchy.html#cfn-quicksight-analysis-columnhierarchy-predefinedhierarchy
	PredefinedHierarchy *Analysis_PredefinedHierarchy `json:"PredefinedHierarchy,omitempty"`

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
func (r *Analysis_ColumnHierarchy) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.ColumnHierarchy"
}
