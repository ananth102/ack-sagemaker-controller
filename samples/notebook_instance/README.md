# Notebook Instance Sample

## Prerequisites

### Common

This sample assumes that you have completed the [common prerequisites](/samples/README.md).

### Update the Notebook Instance Specification

Edit the roleARN value in my-notebook-instance.yaml to include the Sagemaker Execution permissions.

## Using the Notebook Instance operator

### Create a Notebook Instance

This command creates a Sagemaker notebook instance based on the specification provided in my-notebook-instance.yaml.
The Notebook Instance will start at the Pending state and will transition into InService once ready.

```
$ kubectl apply -f my-notebook-instance.yaml
```

### List Notebook Instances
This command lists all the notebook instances created using the ACK controller.
```
$ kubectl get NotebookInstance
```

### Describe a Notebook Instance
This command desribes a specific Notebook Instance, it is useful for checking items like the status, errors or parameters of the Notebook Instance.
```
$ kubectl describe NotebookInstance my-notebook
```

### Delete a Notebook Instance
This command deletes the Notebook Instance.
```
$ kubectl delete NotebookInstance my-notebook
```