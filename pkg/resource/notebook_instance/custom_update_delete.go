package notebook_instance

import (
	"context"
	"fmt"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
)

/*
This function stops the notebook instance(if its running) before the update build request.
*/
func (rm *resourceManager) customPreUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) {

	latestStatus := *latest.ko.Status.NotebookInstanceStatus
	// obj := desired_ptr.RuntimeMetaObject()
	// curr := obj.GetAnnotations()
	curr_annotations := desired.ko.GetAnnotations()
	fmt.Println(" \n \n Checking desired annotations\n", curr_annotations)
	curr_annotations[ackv1alpha1.AnnotationAdopted] = "false"
	desired.ko.SetAnnotations(curr_annotations)
	desired.ko.Annotations["test1ddddd"] = "true"
	// SetMetaDataAnnotation()
	fmt.Println(" \n \n {ttt} Checking desired annotations 2\n", desired.ko.GetAnnotations())

	// curr_latest_annotations := latest.ko.GetAnnotations()
	// fmt.Println(" \n \n Checking latest annotations \n", curr_latest_annotations)
	// curr_latest_annotations["test"] = "TRUE"
	// latest.ko.SetAnnotations(curr_latest_annotations)

	if &latestStatus == nil {
		return
	}
	if latestStatus != svcsdk.NotebookInstanceStatusStopped && latestStatus != svcsdk.NotebookInstanceStatusFailed && latestStatus != svcsdk.NotebookInstanceStatusStopping {
		// if curr == nil {
		// 	curr = make(map[string]string)
		// }
		curr_annotations["stopped_by_ack"] = "True"
		desired.ko.SetAnnotations(curr_annotations)
		// curr["stopped_by_ack"] = "True"
		// obj.SetAnnotations(curr)
		nb_input := svcsdk.StopNotebookInstanceInput{}
		nb_input.NotebookInstanceName = &desired.ko.Name
		rm.sdkapi.StopNotebookInstance(&nb_input)
		fmt.Println(" \n \n Second check \n", desired.ko.GetAnnotations())
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
