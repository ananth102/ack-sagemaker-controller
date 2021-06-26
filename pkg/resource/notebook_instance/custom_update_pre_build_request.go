package notebook_instance

import (
	"context"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
)

func (rm *resourceManager) customUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) {

	nb_input := svcsdk.StopNotebookInstanceInput{}
	nb_input.NotebookInstanceName = &desired.ko.Name
	rm.sdkapi.StopNotebookInstance(&nb_input)
}

func (rm *resourceManager) customDelete(ctx context.Context,
	r *resource) {

	nb_input := svcsdk.StopNotebookInstanceInput{}
	nb_input.NotebookInstanceName = &r.ko.Name
	rm.sdkapi.StopNotebookInstance(&nb_input)
}

// func (rm *resourceManager) customCreate(
// 	ctx context.Context,
// 	r *resource,
// ) {
// 	default_value := "Enabled"
// 	// ko := r.ko.DeepCopy()

// 	if r.ko.Spec.DirectInternetAccess == nil {
// 		r.ko.Spec.DirectInternetAccess = &default_value
// 	}
// 	// r.ko.Spec.
// 	if r.ko.Spec.RootAccess == nil {
// 		r.ko.Spec.RootAccess = &default_value
// 	}
// 	fmt.Println("DIRECT INTERNET  ")
// 	fmt.Println(r.ko.Spec.DirectInternetAccess)
// 	fmt.Println("ROOT")
// 	fmt.Println(r.ko.Spec.RootAccess)
// }
