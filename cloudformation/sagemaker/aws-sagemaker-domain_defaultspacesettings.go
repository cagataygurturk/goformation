// Code generated by "go generate". Please don't change this file directly.

package sagemaker

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Domain_DefaultSpaceSettings AWS CloudFormation Resource (AWS::SageMaker::Domain.DefaultSpaceSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html
type Domain_DefaultSpaceSettings struct {

	// CustomFileSystemConfigs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-customfilesystemconfigs
	CustomFileSystemConfigs []Domain_CustomFileSystemConfig `json:"CustomFileSystemConfigs,omitempty"`

	// CustomPosixUserConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-customposixuserconfig
	CustomPosixUserConfig *Domain_CustomPosixUserConfig `json:"CustomPosixUserConfig,omitempty"`

	// ExecutionRole AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-executionrole
	ExecutionRole string `json:"ExecutionRole"`

	// JupyterLabAppSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-jupyterlabappsettings
	JupyterLabAppSettings *Domain_JupyterLabAppSettings `json:"JupyterLabAppSettings,omitempty"`

	// JupyterServerAppSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-jupyterserverappsettings
	JupyterServerAppSettings *Domain_JupyterServerAppSettings `json:"JupyterServerAppSettings,omitempty"`

	// KernelGatewayAppSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-kernelgatewayappsettings
	KernelGatewayAppSettings *Domain_KernelGatewayAppSettings `json:"KernelGatewayAppSettings,omitempty"`

	// SecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-securitygroups
	SecurityGroups []string `json:"SecurityGroups,omitempty"`

	// SpaceStorageSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-defaultspacesettings.html#cfn-sagemaker-domain-defaultspacesettings-spacestoragesettings
	SpaceStorageSettings *Domain_DefaultSpaceStorageSettings `json:"SpaceStorageSettings,omitempty"`

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
func (r *Domain_DefaultSpaceSettings) AWSCloudFormationType() string {
	return "AWS::SageMaker::Domain.DefaultSpaceSettings"
}
