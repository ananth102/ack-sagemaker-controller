// NotebookInstanceLifecycleConfigName does not get generated by the code generator.
// because the field name in spec vs Describe response do not match
if resp.NotebookInstanceLifecycleConfigName != nil {
	ko.Spec.LifecycleConfigName = resp.NotebookInstanceLifecycleConfigName
} else {
	ko.Spec.LifecycleConfigName = nil
}
err = rm.customSetOutputDescribe(&resource{ko})
if err != nil{
  return nil, err
}