package notebook_instance

import (
	ackcond "github.com/aws-controllers-k8s/runtime/pkg/condition"
	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	corev1 "k8s.io/api/core/v1"
)

func (rm *resourceManager) customSetOutput(
	notebookInstanceStatus *string, ko *svcapitypes.NotebookInstance) {
	if notebookInstanceStatus == nil {
		return
	}

	if *notebookInstanceStatus == svcsdk.NotebookInstanceStatusDeleting || *notebookInstanceStatus == svcsdk.NotebookInstanceStatusFailed || *notebookInstanceStatus == svcsdk.NotebookInstanceStatusInService || *notebookInstanceStatus == svcsdk.NotebookInstanceStatusStopped {
		ackcond.SetSynced(&resource{ko}, corev1.ConditionTrue, nil, nil)

	} else {
		ackcond.SetSynced(&resource{ko}, corev1.ConditionFalse, nil, nil)
	}

}

func (rm *resourceManager) customSetOutputDescribe(r *resource,
	ko *svcapitypes.NotebookInstance) {

	notebook_state := *ko.Status.NotebookInstanceStatus // Get the Notebook State
	rm.customSetOutput(&notebook_state, ko)             // We do this incase the resource is adopted or experienced some other error.

}
