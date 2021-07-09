/* We need this call below in case update succeeds Instantaneously */
if *latest.ko.Status.NotebookInstanceStatus == svcsdk.NotebookInstanceStatusUpdating {
	return nil, requeueWaitWhileUpdating
}