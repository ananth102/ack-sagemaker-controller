package notebook_instance

import (
	"errors"
	"time"

	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
)

var (
	requeueWaitWhileUpdating = ackrequeue.NeededAfter(
		errors.New("Update is in progress."),
		1*time.Second,
	)
)
