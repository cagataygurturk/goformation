// Code generated by "go generate". Please don't change this file directly.

package wisdom

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// KnowledgeBase_ManagedSourceConfiguration AWS CloudFormation Resource (AWS::Wisdom::KnowledgeBase.ManagedSourceConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-managedsourceconfiguration.html
type KnowledgeBase_ManagedSourceConfiguration struct {

	// WebCrawlerConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-knowledgebase-managedsourceconfiguration.html#cfn-wisdom-knowledgebase-managedsourceconfiguration-webcrawlerconfiguration
	WebCrawlerConfiguration *KnowledgeBase_WebCrawlerConfiguration `json:"WebCrawlerConfiguration"`

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
func (r *KnowledgeBase_ManagedSourceConfiguration) AWSCloudFormationType() string {
	return "AWS::Wisdom::KnowledgeBase.ManagedSourceConfiguration"
}
