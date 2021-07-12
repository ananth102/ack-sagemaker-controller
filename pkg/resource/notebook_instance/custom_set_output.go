package notebook_instance

import (
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	corev1 "k8s.io/api/core/v1"
)

func (rm *resourceManager) customSetOutput(
	r *resource,
	notebookInstanceStatus *string,
	ko *svcapitypes.NotebookInstance,
) {
	if notebookInstanceStatus == nil {
		return
	}
	syncConditionStatus := corev1.ConditionUnknown
	if *notebookInstanceStatus == svcsdk.NotebookInstanceStatusDeleting || *notebookInstanceStatus == svcsdk.NotebookInstanceStatusFailed || *notebookInstanceStatus == svcsdk.NotebookInstanceStatusInService || *notebookInstanceStatus == svcsdk.NotebookInstanceStatusStopped {
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
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			condition.Status = syncConditionStatus
			break
		}
	}

}

func (rm *resourceManager) customSetOutputDescribe(r *resource,
	ko *svcapitypes.NotebookInstance) {

	notebook_state := *ko.Status.NotebookInstanceStatus // Get the Notebook State
	/*
		If the notebook is in the stopped state there can be three conditions:
		A. Notebook is stopping for the update - In this case w.Message will be either nothing or  "PRE_UPDATE" and the notebook will update.
		B. The notebook has updated - In this case w.Message will be "POST_UPDATE" and the notebook will start.
		C. The user has stopped the notebook -  In this case w.Message will be "PRE_UPDATE" and the notebook will stay stopped.
	*/
	inServiceSTR := "PRE_UPDATE"
	updatingSTR := "POST_UPDATE"
	if notebook_state == svcsdk.NotebookInstanceStatusStopped {
		for _, w := range ko.Status.Conditions {
			if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
				if *w.Message == updatingSTR {
					val, ok := r.ko.Annotations["stop_after_update"]
					/* If there is an annotation to stop the notebook we will just keep it in the stop state and finish reconciliation. */
					if ok && strings.ToLower(val) == "enabled" {
						for _, w := range ko.Status.Conditions {
							if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
								w.Message = &inServiceSTR
								w.Status = corev1.ConditionTrue //If the user wants the notebook to stop we dont need to reconcile anymore.
								break
							}
						}
					} else {
						/*This code starts the notebook*/
						nb_input := svcsdk.StartNotebookInstanceInput{}
						nb_input.NotebookInstanceName = &r.ko.Name
						rm.sdkapi.StartNotebookInstance(&nb_input)
						for _, w := range ko.Status.Conditions {
							if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
								w.Message = &inServiceSTR
								break
							}

						}
					}
					break
				}
			}
		}

	}
	/*
		This code ensures that the notebook has the proper conditions for updating.
	*/
	if notebook_state == svcsdk.NotebookInstanceStatusPending || notebook_state == svcsdk.NotebookInstanceStatusInService {
		for _, w := range ko.Status.Conditions {
			if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
				w.Message = &inServiceSTR
				if notebook_state == svcsdk.NotebookInstanceStatusInService {
					w.Status = corev1.ConditionTrue //Endpoint is a similar resource that does this.
				}
				break
			}
		}
	}
	rm.customSetOutput(r, &notebook_state, ko) // We do this incase the resource is adopted or experienced some other error.

}
