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

// Code generated by ack-generate. DO NOT EDIT.

package domain

import (
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}

	if ackcompare.HasNilDifference(a.ko.Spec.AppNetworkAccessType, b.ko.Spec.AppNetworkAccessType) {
		delta.Add("Spec.AppNetworkAccessType", a.ko.Spec.AppNetworkAccessType, b.ko.Spec.AppNetworkAccessType)
	} else if a.ko.Spec.AppNetworkAccessType != nil && b.ko.Spec.AppNetworkAccessType != nil {
		if *a.ko.Spec.AppNetworkAccessType != *b.ko.Spec.AppNetworkAccessType {
			delta.Add("Spec.AppNetworkAccessType", a.ko.Spec.AppNetworkAccessType, b.ko.Spec.AppNetworkAccessType)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.AuthMode, b.ko.Spec.AuthMode) {
		delta.Add("Spec.AuthMode", a.ko.Spec.AuthMode, b.ko.Spec.AuthMode)
	} else if a.ko.Spec.AuthMode != nil && b.ko.Spec.AuthMode != nil {
		if *a.ko.Spec.AuthMode != *b.ko.Spec.AuthMode {
			delta.Add("Spec.AuthMode", a.ko.Spec.AuthMode, b.ko.Spec.AuthMode)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings, b.ko.Spec.DefaultUserSettings) {
		delta.Add("Spec.DefaultUserSettings", a.ko.Spec.DefaultUserSettings, b.ko.Spec.DefaultUserSettings)
	} else if a.ko.Spec.DefaultUserSettings != nil && b.ko.Spec.DefaultUserSettings != nil {
		if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.ExecutionRole, b.ko.Spec.DefaultUserSettings.ExecutionRole) {
			delta.Add("Spec.DefaultUserSettings.ExecutionRole", a.ko.Spec.DefaultUserSettings.ExecutionRole, b.ko.Spec.DefaultUserSettings.ExecutionRole)
		} else if a.ko.Spec.DefaultUserSettings.ExecutionRole != nil && b.ko.Spec.DefaultUserSettings.ExecutionRole != nil {
			if *a.ko.Spec.DefaultUserSettings.ExecutionRole != *b.ko.Spec.DefaultUserSettings.ExecutionRole {
				delta.Add("Spec.DefaultUserSettings.ExecutionRole", a.ko.Spec.DefaultUserSettings.ExecutionRole, b.ko.Spec.DefaultUserSettings.ExecutionRole)
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings) {
			delta.Add("Spec.DefaultUserSettings.JupyterServerAppSettings", a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings)
		} else if a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings != nil && b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec) {
				delta.Add("Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec", a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec)
			} else if a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec != nil && b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType) {
					delta.Add("Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType", a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType)
				} else if a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType != nil && b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType != nil {
					if *a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType != *b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType {
						delta.Add("Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType", a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.InstanceType)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN) {
					delta.Add("Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN", a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN)
				} else if a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN != nil && b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN != nil {
					if *a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN != *b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN {
						delta.Add("Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN", a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageARN)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN) {
					delta.Add("Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN", a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN)
				} else if a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN != nil && b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN != nil {
					if *a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN != *b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN {
						delta.Add("Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN", a.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN, b.ko.Spec.DefaultUserSettings.JupyterServerAppSettings.DefaultResourceSpec.SageMakerImageVersionARN)
					}
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings) {
			delta.Add("Spec.DefaultUserSettings.KernelGatewayAppSettings", a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings)
		} else if a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings != nil && b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings != nil {

			if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec) {
				delta.Add("Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec", a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec)
			} else if a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec != nil && b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType) {
					delta.Add("Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType", a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType)
				} else if a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType != nil && b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType != nil {
					if *a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType != *b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType {
						delta.Add("Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType", a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.InstanceType)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN) {
					delta.Add("Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN", a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN)
				} else if a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN != nil && b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN != nil {
					if *a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN != *b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN {
						delta.Add("Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN", a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageARN)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN) {
					delta.Add("Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN", a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN)
				} else if a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN != nil && b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN != nil {
					if *a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN != *b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN {
						delta.Add("Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN", a.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN, b.ko.Spec.DefaultUserSettings.KernelGatewayAppSettings.DefaultResourceSpec.SageMakerImageVersionARN)
					}
				}
			}
		}

		if !ackcompare.SliceStringPEqual(a.ko.Spec.DefaultUserSettings.SecurityGroups, b.ko.Spec.DefaultUserSettings.SecurityGroups) {
			delta.Add("Spec.DefaultUserSettings.SecurityGroups", a.ko.Spec.DefaultUserSettings.SecurityGroups, b.ko.Spec.DefaultUserSettings.SecurityGroups)
		}
		if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.SharingSettings, b.ko.Spec.DefaultUserSettings.SharingSettings) {
			delta.Add("Spec.DefaultUserSettings.SharingSettings", a.ko.Spec.DefaultUserSettings.SharingSettings, b.ko.Spec.DefaultUserSettings.SharingSettings)
		} else if a.ko.Spec.DefaultUserSettings.SharingSettings != nil && b.ko.Spec.DefaultUserSettings.SharingSettings != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption, b.ko.Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption) {
				delta.Add("Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption", a.ko.Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption, b.ko.Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption)
			} else if a.ko.Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption != nil && b.ko.Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption != nil {
				if *a.ko.Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption != *b.ko.Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption {
					delta.Add("Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption", a.ko.Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption, b.ko.Spec.DefaultUserSettings.SharingSettings.NotebookOutputOption)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID, b.ko.Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID) {
				delta.Add("Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID", a.ko.Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID, b.ko.Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID)
			} else if a.ko.Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID != nil && b.ko.Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID != nil {
				if *a.ko.Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID != *b.ko.Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID {
					delta.Add("Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID", a.ko.Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID, b.ko.Spec.DefaultUserSettings.SharingSettings.S3KMSKeyID)
				}
			}
			if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.SharingSettings.S3OutputPath, b.ko.Spec.DefaultUserSettings.SharingSettings.S3OutputPath) {
				delta.Add("Spec.DefaultUserSettings.SharingSettings.S3OutputPath", a.ko.Spec.DefaultUserSettings.SharingSettings.S3OutputPath, b.ko.Spec.DefaultUserSettings.SharingSettings.S3OutputPath)
			} else if a.ko.Spec.DefaultUserSettings.SharingSettings.S3OutputPath != nil && b.ko.Spec.DefaultUserSettings.SharingSettings.S3OutputPath != nil {
				if *a.ko.Spec.DefaultUserSettings.SharingSettings.S3OutputPath != *b.ko.Spec.DefaultUserSettings.SharingSettings.S3OutputPath {
					delta.Add("Spec.DefaultUserSettings.SharingSettings.S3OutputPath", a.ko.Spec.DefaultUserSettings.SharingSettings.S3OutputPath, b.ko.Spec.DefaultUserSettings.SharingSettings.S3OutputPath)
				}
			}
		}
		if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings) {
			delta.Add("Spec.DefaultUserSettings.TensorBoardAppSettings", a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings)
		} else if a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings != nil && b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings != nil {
			if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec) {
				delta.Add("Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec", a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec)
			} else if a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec != nil && b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec != nil {
				if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType) {
					delta.Add("Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType", a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType)
				} else if a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType != nil && b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType != nil {
					if *a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType != *b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType {
						delta.Add("Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType", a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.InstanceType)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN) {
					delta.Add("Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN", a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN)
				} else if a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN != nil && b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN != nil {
					if *a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN != *b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN {
						delta.Add("Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN", a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageARN)
					}
				}
				if ackcompare.HasNilDifference(a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN) {
					delta.Add("Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN", a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN)
				} else if a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN != nil && b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN != nil {
					if *a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN != *b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN {
						delta.Add("Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN", a.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN, b.ko.Spec.DefaultUserSettings.TensorBoardAppSettings.DefaultResourceSpec.SageMakerImageVersionARN)
					}
				}
			}
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DomainID, b.ko.Spec.DomainID) {
		delta.Add("Spec.DomainID", a.ko.Spec.DomainID, b.ko.Spec.DomainID)
	} else if a.ko.Spec.DomainID != nil && b.ko.Spec.DomainID != nil {
		if *a.ko.Spec.DomainID != *b.ko.Spec.DomainID {
			delta.Add("Spec.DomainID", a.ko.Spec.DomainID, b.ko.Spec.DomainID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DomainName, b.ko.Spec.DomainName) {
		delta.Add("Spec.DomainName", a.ko.Spec.DomainName, b.ko.Spec.DomainName)
	} else if a.ko.Spec.DomainName != nil && b.ko.Spec.DomainName != nil {
		if *a.ko.Spec.DomainName != *b.ko.Spec.DomainName {
			delta.Add("Spec.DomainName", a.ko.Spec.DomainName, b.ko.Spec.DomainName)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.HomeEFSFileSystemKMSKeyID, b.ko.Spec.HomeEFSFileSystemKMSKeyID) {
		delta.Add("Spec.HomeEFSFileSystemKMSKeyID", a.ko.Spec.HomeEFSFileSystemKMSKeyID, b.ko.Spec.HomeEFSFileSystemKMSKeyID)
	} else if a.ko.Spec.HomeEFSFileSystemKMSKeyID != nil && b.ko.Spec.HomeEFSFileSystemKMSKeyID != nil {
		if *a.ko.Spec.HomeEFSFileSystemKMSKeyID != *b.ko.Spec.HomeEFSFileSystemKMSKeyID {
			delta.Add("Spec.HomeEFSFileSystemKMSKeyID", a.ko.Spec.HomeEFSFileSystemKMSKeyID, b.ko.Spec.HomeEFSFileSystemKMSKeyID)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID) {
		delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
	} else if a.ko.Spec.KMSKeyID != nil && b.ko.Spec.KMSKeyID != nil {
		if *a.ko.Spec.KMSKeyID != *b.ko.Spec.KMSKeyID {
			delta.Add("Spec.KMSKeyID", a.ko.Spec.KMSKeyID, b.ko.Spec.KMSKeyID)
		}
	}

	if !ackcompare.SliceStringPEqual(a.ko.Spec.SubnetIDs, b.ko.Spec.SubnetIDs) {
		delta.Add("Spec.SubnetIDs", a.ko.Spec.SubnetIDs, b.ko.Spec.SubnetIDs)
	}
	if ackcompare.HasNilDifference(a.ko.Spec.VPCID, b.ko.Spec.VPCID) {
		delta.Add("Spec.VPCID", a.ko.Spec.VPCID, b.ko.Spec.VPCID)
	} else if a.ko.Spec.VPCID != nil && b.ko.Spec.VPCID != nil {
		if *a.ko.Spec.VPCID != *b.ko.Spec.VPCID {
			delta.Add("Spec.VPCID", a.ko.Spec.VPCID, b.ko.Spec.VPCID)
		}
	}

	return delta
}
