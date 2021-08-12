# Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You may
# not use this file except in compliance with the License. A copy of the
# License is located at
#
# 	 http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.
"""Integration tests for the Notebook Lifecycle configuration
"""

import pytest
import logging
import botocore
import datetime

from acktest.k8s import resource as k8s
from acktest.resources import random_suffix_name
from e2e import (
    service_marker,
    wait_for_status,
    create_sagemaker_resource,
    sagemaker_client,
)

from e2e.bootstrap_resources import get_bootstrap_resources
import random

from e2e.replacement_values import REPLACEMENT_VALUES
from time import sleep


@pytest.fixture(scope="function")
def notebook_instance_lifecycleConfig():
    notebook_instance_lfc_name = random_suffix_name("notebookinstancelfc", 32)
    replacements = REPLACEMENT_VALUES.copy()
    replacements["NOTEBOOK_INSTANCE_LFC_NAME"] = notebook_instance_lfc_name
    reference, spec, resource = create_sagemaker_resource(
        resource_plural="notebookinstancelifecycleconfigs",
        resource_name=notebook_instance_lfc_name,
        spec_file="notebook_instance_lifecycle_config",
        replacements=replacements,
    )
    assert resource is not None
    yield (reference, resource, spec)
    if k8s.get_resource_exists(reference):
        _, deleted = k8s.delete_custom_resource(reference, 10, 5)
        assert deleted


def get_notebook_instance_lifecycle_config(notebook_instance_lfc_name: str):
    try:
        resp = sagemaker_client().describe_notebook_instance_lifecycle_config(
            NotebookInstanceLifecycleConfigName=notebook_instance_lfc_name
        )
        return resp
    except botocore.exceptions.ClientError as error:
        logging.error(
            f"SageMaker could not find a Notebook Instance Lifecycle Configuration with the name {notebook_instance_lfc_name}. Error {error}"
        )
        return None


@service_marker
@pytest.mark.canary
class TestNotebookInstanceLifecycleConfig:
    def wait_until_update(self, reference, current_time,wait_period=10,wait_time=5):
        for i in range(wait_period):
            resource = k8s.get_resource(reference)
            assert resource is not None
            assert "lastModifiedTime" in resource["status"]
            lastModifiedTime = resource["status"]["lastModifiedTime"]
            d = datetime.datetime.strptime(lastModifiedTime, "%Y-%m-%dT%H:%M:%SZ")
            if d > current_time:
                return True
            sleep(wait_time)
        return False

    def test_CreateUpdateDeleteNotebookLifecycleConfig(
        self, notebook_instance_lifecycleConfig
    ):
        (reference, resource, spec) = notebook_instance_lifecycleConfig
        assert k8s.get_resource_exists(reference)

        # Getting the resource name
        notebook_instance_lfc_name = resource["spec"].get(
            "notebookInstanceLifecycleConfigName", None
        )
        assert notebook_instance_lfc_name is not None
        current_time = datetime.datetime.today()
        sleep(5)  # Done to avoid flakiness since update happens instantaneously.
        # Verifying that its set correctly
        notebook_instance_lfc_desc = get_notebook_instance_lifecycle_config(
            notebook_instance_lfc_name
        )
        assert (
            notebook_instance_lfc_desc["OnStart"][0]["Content"]
            == spec["spec"]["onStart"][0]["content"]
        )
        assert (
            k8s.get_resource_arn(resource)
            == notebook_instance_lfc_desc["NotebookInstanceLifecycleConfigArn"]
        )
        # We need to keep track of the current time so its best to just do
        # the update test with the create test. update content is pip install six
        update_content = "cGlwIGluc3RhbGwgc2l4"
        spec["spec"]["onStart"] = [
            {"content": update_content}
        ]  # cGlwIGluc3RhbGwgc2l4 = pip install six
        k8s.patch_custom_resource(reference, spec)

        resource = k8s.wait_resource_consumed_by_controller(reference)
        assert resource is not None
        assert (self.wait_until_update(reference, current_time) == True)

        # Verifying that an update was successful
        notebook_instance_lfc_desc = get_notebook_instance_lifecycle_config(
            notebook_instance_lfc_name
        )
        assert (
            notebook_instance_lfc_desc["OnStart"][0]["Content"]
            == update_content
        )

        # Deleting the resource
        _, deleted = k8s.delete_custom_resource(reference, 10, 30)
        assert deleted is True
        assert (
            get_notebook_instance_lifecycle_config(notebook_instance_lfc_name) is None
        )
