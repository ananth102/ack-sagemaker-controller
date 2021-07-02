package notebook_instance_lifecycle_config

import (
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

func customSetDefaults(
	a *resource,
	b *resource,
) {
	if ackcompare.IsNil(a.ko.Spec.OnCreate) && ackcompare.IsNotNil(b.ko.Spec.OnCreate) {
		a.ko.Spec.OnCreate = b.ko.Spec.OnCreate
	}
	if ackcompare.IsNil(a.ko.Spec.OnStart) && ackcompare.IsNotNil(b.ko.Spec.OnStart) {
		a.ko.Spec.OnStart = b.ko.Spec.OnStart
	}
}
