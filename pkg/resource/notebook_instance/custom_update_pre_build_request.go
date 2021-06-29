package notebook_instance

import (
	"context"
	"fmt"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
)

func (rm *resourceManager) customUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) {

	latestStatus := *latest.ko.Status.NotebookInstanceStatus
	fmt.Println("Latest status " + latestStatus)

	// if latestStatus ==

	if &latestStatus == nil {
		return
	}

	if latestStatus != svcsdk.NotebookInstanceStatusStopped && latestStatus != svcsdk.NotebookInstanceStatusFailed && latestStatus != svcsdk.NotebookInstanceStatusStopping {
		nb_input := svcsdk.StopNotebookInstanceInput{}
		nb_input.NotebookInstanceName = &desired.ko.Name
		rm.sdkapi.StopNotebookInstance(&nb_input)
		desired.ko.Annotations["stopped_by_ACK"] = "TRUE"
	} else {
		desired.ko.Annotations["stopped_by_ACK"] = "FALSE"
	}
}

func (rm *resourceManager) customPostUpdate(ctx context.Context,
	desired *resource) {
	// First check if the customer manually stopped the controller
	val, ok := desired.ko.Annotations["stopped_by_ACK"]
	if ok {
		if val == "TRUE" {
			nb_input := svcsdk.StartNotebookInstanceInput{}
			nb_input.NotebookInstanceName = &desired.ko.Name
			rm.sdkapi.StartNotebookInstance(&nb_input)
			desired.ko.Annotations["stopped_by_ACK"] = "FALSE"
		} else {
			return
		}

	} else {
		return
	}

}

func (rm *resourceManager) customCreate(ctx context.Context, r *resource) {

	fmt.Println("ANNOTATIONS")
	fmt.Println(r.ko.Annotations["wowoowowowoowow"])

}

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
