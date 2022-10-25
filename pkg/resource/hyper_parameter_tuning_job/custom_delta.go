// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package hyper_parameter_tuning_job

import (
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

func customSetDefaults(
	a *resource,
	b *resource,
) {
	if ackcompare.IsNotNil(a.ko.Spec.TrainingJobDefinition) && ackcompare.IsNotNil(b.ko.Spec.TrainingJobDefinition) {
		// SageMaker adds StaticHyperParameters prefixed with an underscore. We must ignore these when comparing.
		latestStaticHyperParameters := b.ko.Spec.TrainingJobDefinition.StaticHyperParameters
		if ackcompare.IsNotNil(latestStaticHyperParameters) {
			for key, _ := range latestStaticHyperParameters {
				if key[0:1] == "_" {
					delete(b.ko.Spec.TrainingJobDefinition.StaticHyperParameters, key)
				}
			}
		}
		// TODO: Remove the block below.
		// The server side default of KeepAlivePeriodInSeconds is nil, when launching a HPO job.
		// The code generator currently cannot ignore the field path for resourceConfig.KeepAlivePeriodInSeconds
		// without also ignoring Trainingjob. This block below should be removed once the code generator supports
		// removing fields like resourceConfig.KeepAlivePeriodInSeconds
		if ackcompare.IsNotNil(a.ko.Spec.TrainingJobDefinition) && ackcompare.IsNotNil(b.ko.Spec.TrainingJobDefinition) {
			if ackcompare.IsNotNil(a.ko.Spec.TrainingJobDefinition.ResourceConfig) && ackcompare.IsNotNil(b.ko.Spec.TrainingJobDefinition.ResourceConfig) {
				if ackcompare.IsNotNil(a.ko.Spec.TrainingJobDefinition.ResourceConfig.KeepAlivePeriodInSeconds) && ackcompare.IsNil(b.ko.Spec.TrainingJobDefinition.ResourceConfig.KeepAlivePeriodInSeconds) {
					a.ko.Spec.TrainingJobDefinition.ResourceConfig.KeepAlivePeriodInSeconds = nil
				}
			}
		}
	}
}
