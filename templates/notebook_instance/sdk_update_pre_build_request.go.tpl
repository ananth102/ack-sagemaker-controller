if isNotebookStopping(latest){
    return latest,requeueWaitWhileStopping
}
if isNotebookPending(latest){
    return latest,requeueWaitWhilePending
}
if isNotebookUpdating(latest) && latest.ko.Status.FailureReason == nil {
	return latest, requeueWaitWhileUpdating
}
stopped_by_ack := rm.customPreUpdate(ctx, desired, latest)
if stopped_by_ack {
	curr := latest.ko.GetAnnotations()
	if curr == nil {
		curr = make(map[string]string)
	}
	curr["stopped_by_ack"] = "true"
	latest.ko.SetAnnotations(curr)
	return latest, nil
}