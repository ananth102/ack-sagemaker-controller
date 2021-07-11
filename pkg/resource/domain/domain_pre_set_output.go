package domain

import (
	"strings"

	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
)

func (rm *resourceManager) customDomainPreSetOutput(
	resp *svcsdk.CreateDomainOutput,
	ko *svcapitypes.Domain) {
	if resp.DomainArn != nil {
		arnParts := strings.Split(*resp.DomainArn, "/")
		ko.Spec.DomainID = &arnParts[1]
	}
}
