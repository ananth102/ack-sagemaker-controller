	/*
		This sets resource synced to false so the controller Requeues after it reaches the stopped state.
		If we dont do this we would have to poll Sagemaker once per second.
	*/
	rm.customSetOutputUpdate(desired, ko.Status.NotebookInstanceStatus, ko)
	for _, w := range ko.Status.Conditions {
		if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
			w.Status = corev1.ConditionFalse
			break
		}
		
	}