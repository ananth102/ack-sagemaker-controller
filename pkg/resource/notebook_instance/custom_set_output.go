package notebook_instance

import (
	"context"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	corev1 "k8s.io/api/core/v1"
)

func (rm *resourceManager) customDescribeNotebookSetOutput(ctx context.Context,
	r *resource,
	resp *svcsdk.DescribeNotebookInstanceOutput,
	ko *svcapitypes.NotebookInstance,
) (*svcapitypes.NotebookInstance, error) {
	rm.customSetOutput(r, resp.NotebookInstanceStatus, ko)
	return ko, nil
}

func (rm *resourceManager) customSetOutput(
	r *resource,
	notebookInstanceStatus *string,
	ko *svcapitypes.NotebookInstance,
) {
	if notebookInstanceStatus == nil {
		return
	}
	syncConditionStatus := corev1.ConditionUnknown
	if *notebookInstanceStatus == svcsdk.NotebookInstanceStatusDeleting || *notebookInstanceStatus == svcsdk.NotebookInstanceStatusStopped || *notebookInstanceStatus == svcsdk.NotebookInstanceStatusFailed {
		syncConditionStatus = corev1.ConditionTrue
	} else {
		syncConditionStatus = corev1.ConditionFalse
	}

	var resourceSyncedCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			resourceSyncedCondition = condition
			break
		}
	}
	if resourceSyncedCondition == nil {
		resourceSyncedCondition = &ackv1alpha1.Condition{
			Type: ackv1alpha1.ConditionTypeResourceSynced,
		}
		ko.Status.Conditions = append(ko.Status.Conditions, resourceSyncedCondition)
	}
	resourceSyncedCondition.Status = syncConditionStatus
}
