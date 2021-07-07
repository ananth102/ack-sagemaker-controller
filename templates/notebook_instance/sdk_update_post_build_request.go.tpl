rm.customPostUpdate(ctx, desired, err, latest)
if *latest.ko.Status.NotebookInstanceStatus == svcsdk.NotebookInstanceStatusUpdating {
	return nil, requeueWaitWhileUpdating
}