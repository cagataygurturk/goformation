// Code generated by "go generate". Please don't change this file directly.

package sagemaker

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Domain_StudioWebPortalSettings AWS CloudFormation Resource (AWS::SageMaker::Domain.StudioWebPortalSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-studiowebportalsettings.html
type Domain_StudioWebPortalSettings struct {

	// HiddenAppTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-studiowebportalsettings.html#cfn-sagemaker-domain-studiowebportalsettings-hiddenapptypes
	HiddenAppTypes []string `json:"HiddenAppTypes,omitempty"`

	// HiddenMlTools AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-studiowebportalsettings.html#cfn-sagemaker-domain-studiowebportalsettings-hiddenmltools
	HiddenMlTools []string `json:"HiddenMlTools,omitempty"`

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
func (r *Domain_StudioWebPortalSettings) AWSCloudFormationType() string {
	return "AWS::SageMaker::Domain.StudioWebPortalSettings"
}
