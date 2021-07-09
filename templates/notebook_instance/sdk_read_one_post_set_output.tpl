	 notebook_state := *ko.Status.NotebookInstanceStatus // Get the Notebook State
	/*
	if notebook_state == svcsdk.NotebookInstanceStatusUpdating {
		notebook_annotations["Updating"] = "TRUE"
		r.ko.SetAnnotations(notebook_annotations)
	}
	for _, w := range ko.Status.Conditions {
		fmt.Println("\n \n", notebook_state, " ", w, "\n \n")
	} */
	/*
		If the notebook is in the stopped state there can be three conditions:
		A. Notebook is stopping for the update - In this case ackv1alpha1.ConditionTypeResourceSynced will be true and the notebook will update.
		B. The notebook has updated - In this case ackv1alpha1.ConditionTypeResourceSynced will be false and the notebook will start.
		C. The user has stopped the notebook -  In this case ackv1alpha1.ConditionTypeResourceSynced will be true and the notebook will stay stopped.
	*/
	if notebook_state == svcsdk.NotebookInstanceStatusStopped {
		for _, w := range ko.Status.Conditions {
			if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
				if w.Status == corev1.ConditionFalse {
					/* fmt.Println("\n \n", notebook_state, "   meyooooww   ", w, "\n \n") */
					val, ok := r.ko.Annotations["stop_after_update"]
					/* If there is an annotation to stop the notebook we will just keep it in the stop state and finish reconciliation. */
					if ok && strings.ToLower(val) == "enabled" {
						rm.customSetOutputReadOne(r, aws.String("Stopped"), r.ko)
						/* Finishes reconciliation, code above does not work if ko.Status.Condition is set */
						for _, w := range ko.Status.Conditions {
							if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
								w.Status = corev1.ConditionTrue
								break
							}
							
						}
					} else {
						/*This code starts the notebook and finishes reconciliation*/
						nb_input := svcsdk.StartNotebookInstanceInput{}
						nb_input.NotebookInstanceName = &r.ko.Name
						rm.sdkapi.StartNotebookInstance(&nb_input)
						rm.customSetOutputReadOne(r, aws.String("Pending"), r.ko)
						for _, w := range ko.Status.Conditions {
							if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
								w.Status = corev1.ConditionTrue
								break
							}
							
						}
					}

					break
				}
			}
		}

	}
	/*
		My ec2 instance times out freuqently so I've included this, will take it out for the PR.
		This piece of code performs the same functionality as rm.customSetOutputReadOne and the for loop following that.
	*/
	if notebook_state == svcsdk.NotebookInstanceStatusPending || notebook_state == svcsdk.NotebookInstanceStatusInService {
		for _, w := range ko.Status.Conditions {
			if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
				w.Status = corev1.ConditionTrue
				break
			}
			
		}

	}
	// for _, w := range ko.Status.Conditions {
		/* fmt.Println("\n \n", notebook_state, " PR2 ", w, "\n \n") */
	// }
