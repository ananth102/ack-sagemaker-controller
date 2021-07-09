package notebook_instance

import (
	"errors"

	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
)

var (
	requeueWaitWhileUpdating = ackrequeue.NeededAfter(
		errors.New("Update is in progress."),
		ackrequeue.DefaultRequeueAfterDuration,
	)
	requeueWaitWhileStopped = ackrequeue.NeededAfter(
		errors.New("Update is in progress."),
		ackrequeue.DefaultRequeueAfterDuration,
	)
)
