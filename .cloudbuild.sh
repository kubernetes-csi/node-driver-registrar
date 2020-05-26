#! /bin/bash

. release-tools/prow.sh

# TODO: move the following code into a function in release-tools/prow.sh.

# Register gcloud as a Docker credential helper.
# Required for "docker buildx build --push".
gcloud auth configure-docker

# Extract tag-n-hash value from GIT_TAG (form vYYYYMMDD-tag-n-hash) for REV value.
REV=v$(echo $GIT_TAG | cut -f3- -d 'v')

run_with_go "${CSI_PROW_GO_VERSION_BUILD}" make push-multiarch REV=${REV} REGISTRY_NAME=gcr.io/${STAGING_PROJECT} BUILD_PLATFORMS="${CSI_PROW_BUILD_PLATFORMS}"
