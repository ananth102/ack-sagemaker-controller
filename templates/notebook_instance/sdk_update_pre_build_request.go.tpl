/* This prevents the notebook from finishing reconciliation after it reaches the Updating state */
if *latest.ko.Status.NotebookInstanceStatus == svcsdk.NotebookInstanceStatusUpdating {
		return nil, requeueWaitWhileUpdating
	}
rm.customPreUpdate(ctx, desired, latest, delta)