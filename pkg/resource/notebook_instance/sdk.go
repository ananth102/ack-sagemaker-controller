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

package notebook_instance

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
	_ = &svcapitypes.NotebookInstance{}
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

	var resp *svcsdk.DescribeNotebookInstanceOutput
	resp, err = rm.sdkapi.DescribeNotebookInstanceWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeNotebookInstance", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ValidationException" && strings.HasPrefix(awsErr.Message(), "RecordNotFound") {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.AcceleratorTypes != nil {
		f0 := []*string{}
		for _, f0iter := range resp.AcceleratorTypes {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		ko.Spec.AcceleratorTypes = f0
	} else {
		ko.Spec.AcceleratorTypes = nil
	}
	if resp.AdditionalCodeRepositories != nil {
		f1 := []*string{}
		for _, f1iter := range resp.AdditionalCodeRepositories {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		ko.Spec.AdditionalCodeRepositories = f1
	} else {
		ko.Spec.AdditionalCodeRepositories = nil
	}
	if resp.DefaultCodeRepository != nil {
		ko.Spec.DefaultCodeRepository = resp.DefaultCodeRepository
	} else {
		ko.Spec.DefaultCodeRepository = nil
	}
	if resp.DirectInternetAccess != nil {
		ko.Spec.DirectInternetAccess = resp.DirectInternetAccess
	} else {
		ko.Spec.DirectInternetAccess = nil
	}
	if resp.InstanceType != nil {
		ko.Spec.InstanceType = resp.InstanceType
	} else {
		ko.Spec.InstanceType = nil
	}
	if resp.KmsKeyId != nil {
		ko.Spec.KMSKeyID = resp.KmsKeyId
	} else {
		ko.Spec.KMSKeyID = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.NotebookInstanceArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.NotebookInstanceArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.NotebookInstanceName != nil {
		ko.Spec.NotebookInstanceName = resp.NotebookInstanceName
	} else {
		ko.Spec.NotebookInstanceName = nil
	}
	if resp.NotebookInstanceStatus != nil {
		ko.Status.NotebookInstanceStatus = resp.NotebookInstanceStatus
	} else {
		ko.Status.NotebookInstanceStatus = nil
	}
	if resp.RoleArn != nil {
		ko.Spec.RoleARN = resp.RoleArn
	} else {
		ko.Spec.RoleARN = nil
	}
	if resp.RootAccess != nil {
		ko.Spec.RootAccess = resp.RootAccess
	} else {
		ko.Spec.RootAccess = nil
	}
	if resp.SubnetId != nil {
		ko.Spec.SubnetID = resp.SubnetId
	} else {
		ko.Spec.SubnetID = nil
	}
	if resp.VolumeSizeInGB != nil {
		ko.Spec.VolumeSizeInGB = resp.VolumeSizeInGB
	} else {
		ko.Spec.VolumeSizeInGB = nil
	}

	rm.setStatusDefaults(ko)
	rm.customSetOutputDescribe(r, ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.NotebookInstanceName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeNotebookInstanceInput, error) {
	res := &svcsdk.DescribeNotebookInstanceInput{}

	if r.ko.Spec.NotebookInstanceName != nil {
		res.SetNotebookInstanceName(*r.ko.Spec.NotebookInstanceName)
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

	var resp *svcsdk.CreateNotebookInstanceOutput
	_ = resp
	resp, err = rm.sdkapi.CreateNotebookInstanceWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateNotebookInstance", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.NotebookInstanceArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.NotebookInstanceArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)
	rm.customSetOutput(desired, aws.String(svcsdk.NotebookInstanceStatusPending), ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateNotebookInstanceInput, error) {
	res := &svcsdk.CreateNotebookInstanceInput{}

	if r.ko.Spec.AcceleratorTypes != nil {
		f0 := []*string{}
		for _, f0iter := range r.ko.Spec.AcceleratorTypes {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		res.SetAcceleratorTypes(f0)
	}
	if r.ko.Spec.AdditionalCodeRepositories != nil {
		f1 := []*string{}
		for _, f1iter := range r.ko.Spec.AdditionalCodeRepositories {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		res.SetAdditionalCodeRepositories(f1)
	}
	if r.ko.Spec.DefaultCodeRepository != nil {
		res.SetDefaultCodeRepository(*r.ko.Spec.DefaultCodeRepository)
	}
	if r.ko.Spec.DirectInternetAccess != nil {
		res.SetDirectInternetAccess(*r.ko.Spec.DirectInternetAccess)
	}
	if r.ko.Spec.InstanceType != nil {
		res.SetInstanceType(*r.ko.Spec.InstanceType)
	}
	if r.ko.Spec.KMSKeyID != nil {
		res.SetKmsKeyId(*r.ko.Spec.KMSKeyID)
	}
	if r.ko.Spec.LifecycleConfigName != nil {
		res.SetLifecycleConfigName(*r.ko.Spec.LifecycleConfigName)
	}
	if r.ko.Spec.NotebookInstanceName != nil {
		res.SetNotebookInstanceName(*r.ko.Spec.NotebookInstanceName)
	}
	if r.ko.Spec.RoleARN != nil {
		res.SetRoleArn(*r.ko.Spec.RoleARN)
	}
	if r.ko.Spec.RootAccess != nil {
		res.SetRootAccess(*r.ko.Spec.RootAccess)
	}
	if r.ko.Spec.SecurityGroupIDs != nil {
		f10 := []*string{}
		for _, f10iter := range r.ko.Spec.SecurityGroupIDs {
			var f10elem string
			f10elem = *f10iter
			f10 = append(f10, &f10elem)
		}
		res.SetSecurityGroupIds(f10)
	}
	if r.ko.Spec.SubnetID != nil {
		res.SetSubnetId(*r.ko.Spec.SubnetID)
	}
	if r.ko.Spec.VolumeSizeInGB != nil {
		res.SetVolumeSizeInGB(*r.ko.Spec.VolumeSizeInGB)
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
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer exit(err)
	input, err := rm.newUpdateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateNotebookInstanceOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateNotebookInstanceWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateNotebookInstance", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.UpdateNotebookInstanceInput, error) {
	res := &svcsdk.UpdateNotebookInstanceInput{}

	if r.ko.Spec.AcceleratorTypes != nil {
		f0 := []*string{}
		for _, f0iter := range r.ko.Spec.AcceleratorTypes {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		res.SetAcceleratorTypes(f0)
	}
	if r.ko.Spec.AdditionalCodeRepositories != nil {
		f1 := []*string{}
		for _, f1iter := range r.ko.Spec.AdditionalCodeRepositories {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		res.SetAdditionalCodeRepositories(f1)
	}
	if r.ko.Spec.DefaultCodeRepository != nil {
		res.SetDefaultCodeRepository(*r.ko.Spec.DefaultCodeRepository)
	}
	if r.ko.Spec.InstanceType != nil {
		res.SetInstanceType(*r.ko.Spec.InstanceType)
	}
	if r.ko.Spec.LifecycleConfigName != nil {
		res.SetLifecycleConfigName(*r.ko.Spec.LifecycleConfigName)
	}
	if r.ko.Spec.NotebookInstanceName != nil {
		res.SetNotebookInstanceName(*r.ko.Spec.NotebookInstanceName)
	}
	if r.ko.Spec.RoleARN != nil {
		res.SetRoleArn(*r.ko.Spec.RoleARN)
	}
	if r.ko.Spec.RootAccess != nil {
		res.SetRootAccess(*r.ko.Spec.RootAccess)
	}
	if r.ko.Spec.VolumeSizeInGB != nil {
		res.SetVolumeSizeInGB(*r.ko.Spec.VolumeSizeInGB)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	rm.customDelete(r)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, err = rm.sdkapi.DeleteNotebookInstanceWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteNotebookInstance", err)
	return err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteNotebookInstanceInput, error) {
	res := &svcsdk.DeleteNotebookInstanceInput{}

	if r.ko.Spec.NotebookInstanceName != nil {
		res.SetNotebookInstanceName(*r.ko.Spec.NotebookInstanceName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.NotebookInstance,
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
	case "ResourceLimitExceeded",
		"ResourceNotFound",
		"ResourceInUse",
		"OptInRequired",
		"InvalidParameterCombination",
		"InvalidParameterValue",
		"MissingParameter",
		"MissingAction",
		"InvalidClientTokenId",
		"InvalidQueryParameter",
		"MalformedQueryString",
		"InvalidAction",
		"UnrecognizedClientException",
		"VolumeModificationRateExceeded":
		return true
	default:
		return false
	}
}
