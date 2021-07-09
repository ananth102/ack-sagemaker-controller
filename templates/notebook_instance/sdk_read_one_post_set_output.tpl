	/*
		If the notebook is in the stopped state there can be three conditions:
		A. Notebook is stopping for the update - In this case w.Message will be either nothing or  "DONE_UPDATING" and the notebook will update.
		B. The notebook has updated - In this case w.Message will be "CURRENTLY_UPDATING" and the notebook will start.
		C. The user has stopped the notebook -  In this case w.Message will be "DONE_UPDATING" and the notebook will stay stopped.
	*/
	inServiceSTR := "DONE_UPDATING"
	updatingSTR := "CURRENTLY_UPDATING"
	if notebook_state == svcsdk.NotebookInstanceStatusStopped {
		for _, w := range ko.Status.Conditions {
			if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
				if *w.Message == updatingSTR {
					/* fmt.Println("\n \n", notebook_state, "   meyooooww   ", w, "\n \n") */
					val, ok := r.ko.Annotations["stop_after_update"]
					/* If there is an annotation to stop the notebook we will just keep it in the stop state and finish reconciliation. */
					if ok && strings.ToLower(val) == "enabled" {
						//rm.customSetOutputReadOne(r, aws.String("Stopped"), r.ko)
						/* Finishes reconciliation, code above does not work if ko.Status.Condition is set */
						for _, w := range ko.Status.Conditions {
							if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
								w.Message = &inServiceSTR
								w.Status = corev1.ConditionTrue //If the user wants the notebook to stop we dont need to reconcile anymore.
								break
							}
						}
					} else {
						/*This code starts the notebook*/
						nb_input := svcsdk.StartNotebookInstanceInput{}
						nb_input.NotebookInstanceName = &r.ko.Name
						rm.sdkapi.StartNotebookInstance(&nb_input)
						for _, w := range ko.Status.Conditions {
							if w.Type == ackv1alpha1.ConditionTypeResourceSynced {
								w.Message = &inServiceSTR
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
				w.Message = &inServiceSTR
				if notebook_state == svcsdk.NotebookInstanceStatusInService {
					w.Status = corev1.ConditionTrue //Endpoint is a similar resource that does this.
				}
				break
			}
		}
	}