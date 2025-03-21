// Code generated by "go generate". Please don't change this file directly.

package emr

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Cluster_OnDemandResizingSpecification AWS CloudFormation Resource (AWS::EMR::Cluster.OnDemandResizingSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-ondemandresizingspecification.html
type Cluster_OnDemandResizingSpecification struct {

	// AllocationStrategy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-ondemandresizingspecification.html#cfn-elasticmapreduce-cluster-ondemandresizingspecification-allocationstrategy
	AllocationStrategy *string `json:"AllocationStrategy,omitempty"`

	// CapacityReservationOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-ondemandresizingspecification.html#cfn-elasticmapreduce-cluster-ondemandresizingspecification-capacityreservationoptions
	CapacityReservationOptions *Cluster_OnDemandCapacityReservationOptions `json:"CapacityReservationOptions,omitempty"`

	// TimeoutDurationMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-ondemandresizingspecification.html#cfn-elasticmapreduce-cluster-ondemandresizingspecification-timeoutdurationminutes
	TimeoutDurationMinutes *int `json:"TimeoutDurationMinutes,omitempty"`

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
func (r *Cluster_OnDemandResizingSpecification) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.OnDemandResizingSpecification"
}
