// Code generated by "go generate". Please don't change this file directly.

package entityresolution

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// IdNamespace_IdNamespaceIdMappingWorkflowProperties AWS CloudFormation Resource (AWS::EntityResolution::IdNamespace.IdNamespaceIdMappingWorkflowProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties.html
type IdNamespace_IdNamespaceIdMappingWorkflowProperties struct {

	// IdMappingType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties.html#cfn-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties-idmappingtype
	IdMappingType string `json:"IdMappingType"`

	// ProviderProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties.html#cfn-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties-providerproperties
	ProviderProperties *IdNamespace_NamespaceProviderProperties `json:"ProviderProperties,omitempty"`

	// RuleBasedProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties.html#cfn-entityresolution-idnamespace-idnamespaceidmappingworkflowproperties-rulebasedproperties
	RuleBasedProperties *IdNamespace_NamespaceRuleBasedProperties `json:"RuleBasedProperties,omitempty"`

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
func (r *IdNamespace_IdNamespaceIdMappingWorkflowProperties) AWSCloudFormationType() string {
	return "AWS::EntityResolution::IdNamespace.IdNamespaceIdMappingWorkflowProperties"
}
