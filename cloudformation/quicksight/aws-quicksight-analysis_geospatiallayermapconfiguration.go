// Code generated by "go generate". Please don't change this file directly.

package quicksight

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Analysis_GeospatialLayerMapConfiguration AWS CloudFormation Resource (AWS::QuickSight::Analysis.GeospatialLayerMapConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayermapconfiguration.html
type Analysis_GeospatialLayerMapConfiguration struct {

	// Interactions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayermapconfiguration.html#cfn-quicksight-analysis-geospatiallayermapconfiguration-interactions
	Interactions interface{} `json:"Interactions,omitempty"`

	// Legend AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayermapconfiguration.html#cfn-quicksight-analysis-geospatiallayermapconfiguration-legend
	Legend *Analysis_LegendOptions `json:"Legend,omitempty"`

	// MapLayers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayermapconfiguration.html#cfn-quicksight-analysis-geospatiallayermapconfiguration-maplayers
	MapLayers []Analysis_GeospatialLayerItem `json:"MapLayers,omitempty"`

	// MapState AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayermapconfiguration.html#cfn-quicksight-analysis-geospatiallayermapconfiguration-mapstate
	MapState *Analysis_GeospatialMapState `json:"MapState,omitempty"`

	// MapStyle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-geospatiallayermapconfiguration.html#cfn-quicksight-analysis-geospatiallayermapconfiguration-mapstyle
	MapStyle *Analysis_GeospatialMapStyle `json:"MapStyle,omitempty"`

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
func (r *Analysis_GeospatialLayerMapConfiguration) AWSCloudFormationType() string {
	return "AWS::QuickSight::Analysis.GeospatialLayerMapConfiguration"
}
