package notebook_instance

import (
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

func customSetDefaults(
	a *resource,
	b *resource,
) {
	if ackcompare.IsNil(a.ko.Spec.DirectInternetAccess) && ackcompare.IsNotNil(b.ko.Spec.DirectInternetAccess) {
		a.ko.Spec.DirectInternetAccess = b.ko.Spec.DirectInternetAccess
	}
	if ackcompare.IsNil(a.ko.Spec.RootAccess) && ackcompare.IsNotNil(b.ko.Spec.RootAccess) {
		a.ko.Spec.RootAccess = b.ko.Spec.RootAccess
	}
}
