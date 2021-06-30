package notebook_instance

import (
	"context"

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

	latestStatus := *latest.ko.Status.NotebookInstanceStatus

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

/*
This function starts the notebook instance after the update as long as the annotation desired.ko.Annotations["stopped_by_ACK"] is set to true.
*/

func (rm *resourceManager) customPostUpdate(ctx context.Context,
	desired *resource) {
	val, ok := desired.ko.Annotations["stopped_by_ACK"]
	if ok {
		if val == "TRUE" {
			nb_input := svcsdk.StartNotebookInstanceInput{}
			nb_input.NotebookInstanceName = &desired.ko.Name
			rm.sdkapi.StartNotebookInstance(&nb_input)
			desired.ko.Annotations["stopped_by_ACK"] = "FALSE" //Update cycle is over so we set this to false.
		} else {
			return
		}

	} else {
		//If stopped_by_ACK does not even exist that means the controller did not stop the notebook
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
