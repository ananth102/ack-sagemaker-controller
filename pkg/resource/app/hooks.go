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

package app

import (
	"context"
	"errors"

	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	svccommon "github.com/aws-controllers-k8s/sagemaker-controller/pkg/common"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
)

var (
	modifyingStatuses = []string{svcsdk.AppStatusPending,
		svcsdk.AppStatusDeleting}

	resourceName = resourceGK.Kind

	requeueWaitWhileDeleting = ackrequeue.NeededAfter(
		errors.New(resourceName+" is deleting."),
		ackrequeue.DefaultRequeueAfterDuration,
	)
)

// requeueUntilCanModify creates and returns an
// ackrequeue error if a resource's latest status matches
// any of the defined modifying statuses below.
func (rm *resourceManager) requeueUntilCanModify(
	ctx context.Context,
	r *resource,
) error {
	latestStatus := r.ko.Status.AppStatus
	return svccommon.RequeueIfModifying(latestStatus, &resourceName, &modifyingStatuses)
}

// customSetOutput sets the ack syncedCondition depending on
// whether the latest status of the resource is one of the
// defined modifyingStatuses.
func (rm *resourceManager) customSetOutput(r *resource) {
	latestStatus := r.ko.Status.AppStatus
	svccommon.SetSyncedCondition(r, latestStatus, &resourceName, &modifyingStatuses)
}

// customSetOutput sets the ack syncedCondition depending on
// whether the latest status of the resource is one of the
// defined modifyingStatuses.
func (rm *resourceManager) customSetCreate(r *resource) {
	svccommon.SetSyncedCondition(r, aws.String(svcsdk.AppStatusPending), &resourceName, &modifyingStatuses)
}
