// Code generated by "go generate". Please don't change this file directly.

package notificationscontacts

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// EmailContact_EmailContact AWS CloudFormation Resource (AWS::NotificationsContacts::EmailContact.EmailContact)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html
type EmailContact_EmailContact struct {

	// Address AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-address
	Address string `json:"Address"`

	// Arn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-arn
	Arn string `json:"Arn"`

	// CreationTime AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-creationtime
	CreationTime string `json:"CreationTime"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-name
	Name string `json:"Name"`

	// Status AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-status
	Status string `json:"Status"`

	// UpdateTime AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-notificationscontacts-emailcontact-emailcontact.html#cfn-notificationscontacts-emailcontact-emailcontact-updatetime
	UpdateTime string `json:"UpdateTime"`

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
func (r *EmailContact_EmailContact) AWSCloudFormationType() string {
	return "AWS::NotificationsContacts::EmailContact.EmailContact"
}
