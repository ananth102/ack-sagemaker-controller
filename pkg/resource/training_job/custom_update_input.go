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

// Use this file if conditions need to be updated based on the latest status
// of training job which is not evident from API response

package training_job

import (
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
)

// Because a rule configuration and a profiler configuration is needed for profiler
// it is fair to assume that d
func disableProfilerCheck(desired *resource, latest *resource) bool {
	if ackcompare.IsNotNil(desired.ko.Spec) && ackcompare.IsNotNil(latest.ko.Spec) {
		if ackcompare.IsNil(desired.ko.Spec.ProfilerRuleConfigurations) && ackcompare.IsNotNil(latest.ko.Spec.ProfilerRuleConfigurations) {
			return true
		}
		if ackcompare.IsNil(desired.ko.Spec.ProfilerConfig) && ackcompare.IsNotNil(latest.ko.Spec.ProfilerConfig) {
			return true
		}
	}
	return false
}

func customSetDisableProfiler(updateInput *svcsdk.UpdateTrainingJobInput) {
	if ackcompare.IsNil(updateInput.ProfilerConfig) {
		profilerConfigUpdateInput := &svcsdk.ProfilerConfigForUpdate{}
		profilerConfigUpdateInput.SetDisableProfiler(true)
		updateInput.SetProfilerConfig(profilerConfigUpdateInput)
	} else {
		updateInput.ProfilerConfig.SetDisableProfiler(true)
	}

}
