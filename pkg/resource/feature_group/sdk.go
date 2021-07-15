// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package feature_group

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.SageMaker{}
	_ = &svcapitypes.FeatureGroup{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeFeatureGroupOutput
	resp, err = rm.sdkapi.DescribeFeatureGroupWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeFeatureGroup", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ResourceNotFound" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Description != nil {
		ko.Spec.Description = resp.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.EventTimeFeatureName != nil {
		ko.Spec.EventTimeFeatureName = resp.EventTimeFeatureName
	} else {
		ko.Spec.EventTimeFeatureName = nil
	}
	if resp.FailureReason != nil {
		ko.Status.FailureReason = resp.FailureReason
	} else {
		ko.Status.FailureReason = nil
	}
	if resp.FeatureDefinitions != nil {
		f4 := []*svcapitypes.FeatureDefinition{}
		for _, f4iter := range resp.FeatureDefinitions {
			f4elem := &svcapitypes.FeatureDefinition{}
			if f4iter.FeatureName != nil {
				f4elem.FeatureName = f4iter.FeatureName
			}
			if f4iter.FeatureType != nil {
				f4elem.FeatureType = f4iter.FeatureType
			}
			f4 = append(f4, f4elem)
		}
		ko.Spec.FeatureDefinitions = f4
	} else {
		ko.Spec.FeatureDefinitions = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.FeatureGroupArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.FeatureGroupArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.FeatureGroupName != nil {
		ko.Spec.FeatureGroupName = resp.FeatureGroupName
	} else {
		ko.Spec.FeatureGroupName = nil
	}
	if resp.FeatureGroupStatus != nil {
		ko.Status.FeatureGroupStatus = resp.FeatureGroupStatus
	} else {
		ko.Status.FeatureGroupStatus = nil
	}
	if resp.OfflineStoreConfig != nil {
		f9 := &svcapitypes.OfflineStoreConfig{}
		if resp.OfflineStoreConfig.DataCatalogConfig != nil {
			f9f0 := &svcapitypes.DataCatalogConfig{}
			if resp.OfflineStoreConfig.DataCatalogConfig.Catalog != nil {
				f9f0.Catalog = resp.OfflineStoreConfig.DataCatalogConfig.Catalog
			}
			if resp.OfflineStoreConfig.DataCatalogConfig.Database != nil {
				f9f0.Database = resp.OfflineStoreConfig.DataCatalogConfig.Database
			}
			if resp.OfflineStoreConfig.DataCatalogConfig.TableName != nil {
				f9f0.TableName = resp.OfflineStoreConfig.DataCatalogConfig.TableName
			}
			f9.DataCatalogConfig = f9f0
		}
		if resp.OfflineStoreConfig.DisableGlueTableCreation != nil {
			f9.DisableGlueTableCreation = resp.OfflineStoreConfig.DisableGlueTableCreation
		}
		if resp.OfflineStoreConfig.S3StorageConfig != nil {
			f9f2 := &svcapitypes.S3StorageConfig{}
			if resp.OfflineStoreConfig.S3StorageConfig.KmsKeyId != nil {
				f9f2.KMSKeyID = resp.OfflineStoreConfig.S3StorageConfig.KmsKeyId
			}
			if resp.OfflineStoreConfig.S3StorageConfig.ResolvedOutputS3Uri != nil {
				f9f2.ResolvedOutputS3URI = resp.OfflineStoreConfig.S3StorageConfig.ResolvedOutputS3Uri
			}
			if resp.OfflineStoreConfig.S3StorageConfig.S3Uri != nil {
				f9f2.S3URI = resp.OfflineStoreConfig.S3StorageConfig.S3Uri
			}
			f9.S3StorageConfig = f9f2
		}
		ko.Spec.OfflineStoreConfig = f9
	} else {
		ko.Spec.OfflineStoreConfig = nil
	}
	if resp.OnlineStoreConfig != nil {
		f11 := &svcapitypes.OnlineStoreConfig{}
		if resp.OnlineStoreConfig.EnableOnlineStore != nil {
			f11.EnableOnlineStore = resp.OnlineStoreConfig.EnableOnlineStore
		}
		if resp.OnlineStoreConfig.SecurityConfig != nil {
			f11f1 := &svcapitypes.OnlineStoreSecurityConfig{}
			if resp.OnlineStoreConfig.SecurityConfig.KmsKeyId != nil {
				f11f1.KMSKeyID = resp.OnlineStoreConfig.SecurityConfig.KmsKeyId
			}
			f11.SecurityConfig = f11f1
		}
		ko.Spec.OnlineStoreConfig = f11
	} else {
		ko.Spec.OnlineStoreConfig = nil
	}
	if resp.RecordIdentifierFeatureName != nil {
		ko.Spec.RecordIdentifierFeatureName = resp.RecordIdentifierFeatureName
	} else {
		ko.Spec.RecordIdentifierFeatureName = nil
	}
	if resp.RoleArn != nil {
		ko.Spec.RoleARN = resp.RoleArn
	} else {
		ko.Spec.RoleARN = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.FeatureGroupName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeFeatureGroupInput, error) {
	res := &svcsdk.DescribeFeatureGroupInput{}

	if r.ko.Spec.FeatureGroupName != nil {
		res.SetFeatureGroupName(*r.ko.Spec.FeatureGroupName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateFeatureGroupOutput
	_ = resp
	resp, err = rm.sdkapi.CreateFeatureGroupWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateFeatureGroup", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.FeatureGroupArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.FeatureGroupArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateFeatureGroupInput, error) {
	res := &svcsdk.CreateFeatureGroupInput{}

	if r.ko.Spec.Description != nil {
		res.SetDescription(*r.ko.Spec.Description)
	}
	if r.ko.Spec.EventTimeFeatureName != nil {
		res.SetEventTimeFeatureName(*r.ko.Spec.EventTimeFeatureName)
	}
	if r.ko.Spec.FeatureDefinitions != nil {
		f2 := []*svcsdk.FeatureDefinition{}
		for _, f2iter := range r.ko.Spec.FeatureDefinitions {
			f2elem := &svcsdk.FeatureDefinition{}
			if f2iter.FeatureName != nil {
				f2elem.SetFeatureName(*f2iter.FeatureName)
			}
			if f2iter.FeatureType != nil {
				f2elem.SetFeatureType(*f2iter.FeatureType)
			}
			f2 = append(f2, f2elem)
		}
		res.SetFeatureDefinitions(f2)
	}
	if r.ko.Spec.FeatureGroupName != nil {
		res.SetFeatureGroupName(*r.ko.Spec.FeatureGroupName)
	}
	if r.ko.Spec.OfflineStoreConfig != nil {
		f4 := &svcsdk.OfflineStoreConfig{}
		if r.ko.Spec.OfflineStoreConfig.DataCatalogConfig != nil {
			f4f0 := &svcsdk.DataCatalogConfig{}
			if r.ko.Spec.OfflineStoreConfig.DataCatalogConfig.Catalog != nil {
				f4f0.SetCatalog(*r.ko.Spec.OfflineStoreConfig.DataCatalogConfig.Catalog)
			}
			if r.ko.Spec.OfflineStoreConfig.DataCatalogConfig.Database != nil {
				f4f0.SetDatabase(*r.ko.Spec.OfflineStoreConfig.DataCatalogConfig.Database)
			}
			if r.ko.Spec.OfflineStoreConfig.DataCatalogConfig.TableName != nil {
				f4f0.SetTableName(*r.ko.Spec.OfflineStoreConfig.DataCatalogConfig.TableName)
			}
			f4.SetDataCatalogConfig(f4f0)
		}
		if r.ko.Spec.OfflineStoreConfig.DisableGlueTableCreation != nil {
			f4.SetDisableGlueTableCreation(*r.ko.Spec.OfflineStoreConfig.DisableGlueTableCreation)
		}
		if r.ko.Spec.OfflineStoreConfig.S3StorageConfig != nil {
			f4f2 := &svcsdk.S3StorageConfig{}
			if r.ko.Spec.OfflineStoreConfig.S3StorageConfig.KMSKeyID != nil {
				f4f2.SetKmsKeyId(*r.ko.Spec.OfflineStoreConfig.S3StorageConfig.KMSKeyID)
			}
			if r.ko.Spec.OfflineStoreConfig.S3StorageConfig.ResolvedOutputS3URI != nil {
				f4f2.SetResolvedOutputS3Uri(*r.ko.Spec.OfflineStoreConfig.S3StorageConfig.ResolvedOutputS3URI)
			}
			if r.ko.Spec.OfflineStoreConfig.S3StorageConfig.S3URI != nil {
				f4f2.SetS3Uri(*r.ko.Spec.OfflineStoreConfig.S3StorageConfig.S3URI)
			}
			f4.SetS3StorageConfig(f4f2)
		}
		res.SetOfflineStoreConfig(f4)
	}
	if r.ko.Spec.OnlineStoreConfig != nil {
		f5 := &svcsdk.OnlineStoreConfig{}
		if r.ko.Spec.OnlineStoreConfig.EnableOnlineStore != nil {
			f5.SetEnableOnlineStore(*r.ko.Spec.OnlineStoreConfig.EnableOnlineStore)
		}
		if r.ko.Spec.OnlineStoreConfig.SecurityConfig != nil {
			f5f1 := &svcsdk.OnlineStoreSecurityConfig{}
			if r.ko.Spec.OnlineStoreConfig.SecurityConfig.KMSKeyID != nil {
				f5f1.SetKmsKeyId(*r.ko.Spec.OnlineStoreConfig.SecurityConfig.KMSKeyID)
			}
			f5.SetSecurityConfig(f5f1)
		}
		res.SetOnlineStoreConfig(f5)
	}
	if r.ko.Spec.RecordIdentifierFeatureName != nil {
		res.SetRecordIdentifierFeatureName(*r.ko.Spec.RecordIdentifierFeatureName)
	}
	if r.ko.Spec.RoleARN != nil {
		res.SetRoleArn(*r.ko.Spec.RoleARN)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, err = rm.sdkapi.DeleteFeatureGroupWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteFeatureGroup", err)
	return err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteFeatureGroupInput, error) {
	res := &svcsdk.DeleteFeatureGroupInput{}

	if r.ko.Spec.FeatureGroupName != nil {
		res.SetFeatureGroupName(*r.ko.Spec.FeatureGroupName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.FeatureGroup,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Message()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "ResourceInUse",
		"ResourceLimitExceeded",
		"ResourceNotFound",
		"InvalidAction",
		"InvalidClientTokenId",
		"InvalidParameterCombination",
		"InvalidParameterValue",
		"InvalidQueryParameter",
		"MalformedQueryString",
		"MissingAction",
		"MissingParameter",
		"OptInRequired":
		return true
	default:
		return false
	}
}
