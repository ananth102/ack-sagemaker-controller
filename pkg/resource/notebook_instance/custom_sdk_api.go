package notebook_instance

import (
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
)

func (rm *resourceManager) stopNotebookInstance(r *resource) error {
	nb_input := svcsdk.StopNotebookInstanceInput{}
	nb_input.NotebookInstanceName = &r.ko.Name
	_, err := rm.sdkapi.StopNotebookInstance(&nb_input) //Stop Notebook Instance does not return a response
	rm.metrics.RecordAPICall("STOP", "StopNotebookInstance", err)
	return err
}

func (rm *resourceManager) startNotebookInstance(r *resource) {
	nb_input := svcsdk.StartNotebookInstanceInput{}
	nb_input.NotebookInstanceName = &r.ko.Name
	_, err := rm.sdkapi.StartNotebookInstance(&nb_input) //Start Notebook Instance does not return a response
	rm.metrics.RecordAPICall("START", "StartNotebookInstance", err)
}
