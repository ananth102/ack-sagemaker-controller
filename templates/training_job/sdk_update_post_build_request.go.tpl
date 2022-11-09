warmpool_diff := delta.DifferentAt("Spec.ResourceConfig.KeepAlivePeriodInSeconds")
profiler_diff := delta.DifferentAt("Spec.ProfilerConfig") && delta.DifferentAt("Spec.ProfilerRuleConfigurations")
trainingJobStatus := latest.ko.Status.TrainingJobStatus
if warmpool_diff && profiler_diff{
	return latest, nil
}
if warmpool_diff {
	warmpool_terminal := customWarmPool(latest)
	if warmpool_terminal {
		return latest, nil
	}
	if err := customSetOutputUpdate(latest); err != nil {
		return nil,err
	}
}
if profiler_diff {
	input.SetResourceConfig(nil)
	if trainingJobStatus != nil && *trainingJobStatus != svcsdk.TrainingJobStatusInProgress{
		return latest, nil
	}
	if disableProfilerCheck(desired, latest) {
		customSetDisableProfiler(input)
	}
}