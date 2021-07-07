package notebook_instance

import (
	"context"
	"fmt"
	"strings"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
)

/*
This function stops the notebook instance(if its running) before the update build request.
It also keeps track of whether the notebook was stopped beforehand.
*/
func (rm *resourceManager) customUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) {

	// delta.

	desired_ptr := rm.concreteResource(desired)

	latestStatus := *latest.ko.Status.NotebookInstanceStatus
	obj := desired_ptr.RuntimeMetaObject()
	curr := obj.GetAnnotations()
	fmt.Println(" \n \n First check \n", curr)
	if &latestStatus == nil {
		return
	}
	if latestStatus != svcsdk.NotebookInstanceStatusStopped && latestStatus != svcsdk.NotebookInstanceStatusFailed && latestStatus != svcsdk.NotebookInstanceStatusStopping {
		if curr == nil {
			curr = make(map[string]string)
		}
		curr["stopped_by_ack"] = "True"
		obj.SetAnnotations(curr)
		nb_input := svcsdk.StopNotebookInstanceInput{}
		nb_input.NotebookInstanceName = &desired.ko.Name
		rm.sdkapi.StopNotebookInstance(&nb_input)
		fmt.Println(" \n \n Second check \n", curr)
	}
}

/*
This function starts the notebook instance after the update as long as the annotation desired.ko.Annotations["stopped_by_ACK"] is set to true.
*/

func (rm *resourceManager) customPostUpdate(ctx context.Context,
	desired *resource, err error, latest *resource) {

	if err != nil && *latest.ko.Status.NotebookInstanceStatus != svcsdk.NotebookInstanceStatusUpdating {
		return
	}

	val, ok := desired.ko.Annotations["stop_after_update"]

	nb_input := svcsdk.StartNotebookInstanceInput{}
	nb_input.NotebookInstanceName = &desired.ko.Name
	fmt.Println("\n  PRINT WORKSs  ", ok, " \n \n")
	if ok {
		santizedStop := strings.ToLower(val)
		if santizedStop != "enabled" {
			nb_input := svcsdk.StartNotebookInstanceInput{}
			nb_input.NotebookInstanceName = &desired.ko.Name
			rm.sdkapi.StartNotebookInstance(&nb_input)
		} else {
			return
		}

	} else {
		fmt.Println("\n \n INSIDE HERE \n \n")
		nb_input := svcsdk.StartNotebookInstanceInput{}
		nb_input.NotebookInstanceName = &desired.ko.Name
		rm.sdkapi.StartNotebookInstance(&nb_input)
		return
	}

}

/*
This code stops the NotebookInstance right before its about to be deleted.
*/
func (rm *resourceManager) customDelete(ctx context.Context,
	r *resource) {

	latestStatus := *r.ko.Status.NotebookInstanceStatus

	if &latestStatus == nil {
		return
	}

	//We only want to stop the Notebook if its not already stopped/stopping or is in a failed state.
	if latestStatus != svcsdk.NotebookInstanceStatusStopped && latestStatus != svcsdk.NotebookInstanceStatusFailed && latestStatus != svcsdk.NotebookInstanceStatusStopping {
		nb_input := svcsdk.StopNotebookInstanceInput{}
		nb_input.NotebookInstanceName = &r.ko.Name
		rm.sdkapi.StopNotebookInstance(&nb_input)
	}
}

// func (rm *resourceManager) CustomUpdateConditions(
// 	ko *svcapitypes.NotebookInstance,
// 	r *resource,
// 	err error,
// ) bool {

// 	//First check if the annotation exists
// 	//If it does exist and is not enabled and notebook is not in pending or inservice return false
// 	//otherwise return true
// 	val, ok := r.ko.Annotations["stop_after_update"]
// 	if ok && val != "enabled" && !(*ko.Status.NotebookInstanceStatus != svcsdk.NotebookInstanceStatusPending && *ko.Status.NotebookInstanceStatus != svcsdk.NotebookInstanceStatusInService) {
// 		return false
// 	}
// 	if !ok {
// 		if !(*ko.Status.NotebookInstanceStatus != svcsdk.NotebookInstanceStatusPending && *ko.Status.NotebookInstanceStatus != svcsdk.NotebookInstanceStatusInService) {
// 			return false
// 		}
// 	}
// 	return true

// }
