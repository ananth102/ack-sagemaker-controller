if *latest.ko.Status.NotebookInstanceStatus == svcsdk.NotebookInstanceStatusUpdating {
		rm.customPostUpdate(ctx, desired, err, latest)
		return nil, requeueWaitWhileUpdating
	}
rm.customUpdate(ctx, desired, latest, delta)
