# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

.PHONY: push-multiarch-% push-multiarch

CMDS=csi-node-driver-registrar
all: build

include release-tools/build.make

# Additional parameters are needed when pushing to a local registry,
# see https://github.com/docker/buildx/issues/94.
# However, that then runs into https://github.com/docker/cli/issues/2396.
#
# What works for local testing is:
# make push-multiarch PULL_BASE_REF=master REGISTRY_NAME=<your account on dockerhub.io> BUILD_PLATFORMS="linux amd64; windows amd64 .exe; linux ppc64le -ppc64le; linux s390x -s390x"
DOCKER_BUILDX_CREATE_ARGS ?=

# This target builds a multiarch image for one command using Moby BuildKit builder toolkit.
# Docker Buildx is included in Docker 19.03.
#
# ./cmd/<command>/Dockerfile[.Windows] is used if found, otherwise Dockerfile[.Windows].
# BUILD_PLATFORMS determines which individual images are included in the multiarch image.
# PULL_BASE_REF must be set to 'master', 'release-x.y', or a tag name, and determines
# the tag for the resulting multiarch image.
push-multiarch-%: check-pull-base-ref build-%
	set -ex; \
	DOCKER_CLI_EXPERIMENTAL=enabled; \
	export DOCKER_CLI_EXPERIMENTAL; \
	docker buildx create $(DOCKER_BUILDX_CREATE_ARGS) --use --name multiarchimage-buildertest; \
	trap "docker buildx rm multiarchimage-buildertest" EXIT; \
	dockerfile_linux=$$(if [ -e ./cmd/$*/Dockerfile ]; then echo ./cmd/$*/Dockerfile; else echo Dockerfile; fi); \
	dockerfile_windows=$$(if [ -e ./cmd/$*/Dockerfile.Windows ]; then echo ./cmd/$*/Dockerfile.Windows; else echo Dockerfile.Windows; fi); \
	build_platforms='$(BUILD_PLATFORMS)'; \
	if ! [ "$$build_platforms" ]; then build_platforms="linux amd64"; fi; \
	pushMultiArch () { \
		tag=$$1; \
		echo "$$build_platforms" | tr ';' '\n' | while read -r os arch suffix; do \
			docker buildx build --push \
				--tag $(IMAGE_NAME):$$arch-$$os-$$tag \
				--platform=$$os/$$arch \
				--file $$(eval echo \$${dockerfile_$$os}) \
				--build-arg binary=./bin/$*$$suffix \
				--label revision=$(REV) \
				.; \
		done; \
		images=$$(echo "$$build_platforms" | tr ';' '\n' | while read -r os arch suffix; do echo $(IMAGE_NAME):$$arch-$$os-$$tag; done); \
		docker manifest create --amend $(IMAGE_NAME):$$tag $$images; \
		docker manifest push -p $(IMAGE_NAME):$$tag; \
	}; \
	if [ $(PULL_BASE_REF) = "master" ]; then \
		       : "creating or overwriting canary image"; \
		       pushMultiArch canary; \
	elif echo $(PULL_BASE_REF) | grep -q -e 'release-*' ; then \
		       : "creating or overwriting canary image for release branch"; \
			release_canary_tag=$$(echo $(PULL_BASE_REF) | cut -f2 -d '-')-canary; \
			pushMultiArch $$release_canary_tag; \
	elif docker pull $(IMAGE_NAME):$(PULL_BASE_REF) 2>&1 | tee /dev/stderr | grep -q "manifest for $(IMAGE_NAME):$(PULL_BASE_REF) not found"; then \
		       : "creating release image"; \
		       pushMultiArch $(PULL_BASE_REF); \
	else \
		       : "release image $(IMAGE_NAME):$(PULL_BASE_REF) already exists, skipping push"; \
	fi; \

.PHONY: check-pull-base-ref
check-pull-base-ref:
	if ! [ "$(PULL_BASE_REF)" ]; then \
		echo >&2 "ERROR: PULL_BASE_REF must be set to 'master', 'release-x.y', or a tag name."; \
		exit 1; \
	fi

push-multiarch: $(CMDS:%=push-multiarch-%)
