import boto3
import botocore
import pytest
import logging
from typing import Dict

from acktest.resources import random_suffix_name
from acktest.k8s import resource as k8s

from e2e import (
    service_marker,
    create_sagemaker_resource,
    wait_for_status,
    sagemaker_client,
    assert_tags_in_sync,
)
from e2e.replacement_values import REPLACEMENT_VALUES


def createExample():
    res = client.create_domain(DomainName="my-dom",AuthMode='IAM',DefaultUserSettings={
        "ExecutionRole" : SAGEMAKER_EXECUTION_ROLE_ARN
    },
    SubnetIds=["subnet-278aae6b","subnet-fa9c1a91","subnet-7437cb09"],
    VpcId="vpc-6aec7201"
    )
    domainId = res["DomainArn"].split("/")[1]
    return domainId