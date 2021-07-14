package notebook_instance

import (
	"errors"
	"time"

	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
)

var (
	requeueWaitWhileStopping = ackrequeue.NeededAfter(
		errors.New("NotebookInstance in 'Stopping' state, cannot be modified or deleted"),
		10*time.Second,
	)
	requeueWaitWhilePending = ackrequeue.NeededAfter(
		errors.New("NotebookInstance in 'Pending' state, cannot be modified or deleted"),
		10*time.Second,
	)
)

func isNotebookStopping(r *resource) bool {
	if r.ko.Status.NotebookInstanceStatus == nil {
		return false
	}
	notebookInstanceStatus := r.ko.Status.NotebookInstanceStatus

	return *notebookInstanceStatus == svcsdk.NotebookInstanceStatusStopping
}

func isNotebookPending(r *resource) bool {
	if r.ko.Status.NotebookInstanceStatus == nil {
		return false
	}
	notebookInstanceStatus := r.ko.Status.NotebookInstanceStatus

	return *notebookInstanceStatus == svcsdk.NotebookInstanceStatusPending
}
