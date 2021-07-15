package notebook_instance_lifecycle_config

import (
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

func modifyDeltaCreate(
	delta *ackcompare.Delta,
	a *resource,
	b *resource,
) *ackcompare.Delta {
	if ackcompare.HasNilDifference(a.ko.Spec.OnCreate, b.ko.Spec.OnCreate) {
		delta.Add("Spec.OnCreate", a.ko.Spec.OnCreate, b.ko.Spec.OnCreate)
		return delta
	}
	//check length
	if a.ko.Spec.OnCreate != nil && b.ko.Spec.OnCreate != nil {
		if len(a.ko.Spec.OnCreate) != len(b.ko.Spec.OnCreate) {
			delta.Add("Spec.OnCreate", a.ko.Spec.OnCreate, b.ko.Spec.OnCreate)
			return delta
		}
	} else {
		return delta
	}
	//Check variables, a and b have to be of equal length
	onCreateLen := len(a.ko.Spec.OnCreate)
	if a.ko.Spec.OnCreate != nil && b.ko.Spec.OnCreate != nil {
		for i := 0; i < onCreateLen; i++ {
			if a.ko.Spec.OnCreate[i].Content != b.ko.Spec.OnCreate[i].Content {
				delta.Add("Spec.OnCreate", a.ko.Spec.OnCreate, b.ko.Spec.OnCreate)
				return delta
			}
		}
	}
	return delta
}

func modifyDeltaStart(
	delta *ackcompare.Delta,
	a *resource,
	b *resource,
) *ackcompare.Delta {
	if ackcompare.HasNilDifference(a.ko.Spec.OnStart, b.ko.Spec.OnStart) {
		delta.Add("Spec.OnStart", a.ko.Spec.OnStart, b.ko.Spec.OnStart)
		return delta
	}
	//check length
	if a.ko.Spec.OnStart != nil && b.ko.Spec.OnStart != nil {
		if len(a.ko.Spec.OnStart) != len(b.ko.Spec.OnStart) {
			delta.Add("Spec.OnStart", a.ko.Spec.OnStart, b.ko.Spec.OnStart)
			return delta
		}
	} else {
		return delta
	}
	//Check variables, a and b have to be of equal length
	onStartLen := len(a.ko.Spec.OnStart)
	if a.ko.Spec.OnStart != nil && b.ko.Spec.OnStart != nil {
		for i := 0; i < onStartLen; i++ {
			if a.ko.Spec.OnStart[i].Content != b.ko.Spec.OnStart[i].Content {
				delta.Add("Spec.OnStart", a.ko.Spec.OnStart, b.ko.Spec.OnStart)
				return delta
			}
		}
	}
	return delta
}
