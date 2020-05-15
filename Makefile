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

# This target builds multiarch images using Moby BuildKit builder toolkit.
# Docker Buildx is included in Docker 19.03 and needs DOCKER_CLI_EXPERIMENTAL enabled to run corresponding commands.
# Currently amd, s390x and Windows manifest is pushed for canary, release branch and released tags.
# Images generated from Prow build are pushed to staging area on gcr
push-multiarch-%:
	make BUILD_PLATFORMS="windows amd64 .exe"
	set -ex; \
	gcloud auth configure-docker; \
	DOCKER_CLI_EXPERIMENTAL=enabled; \
	export DOCKER_CLI_EXPERIMENTAL; \
	docker buildx create --use --name multiarchimage-buildertest; \
	pushMultiArch () { \
                tag=$$1; \
                docker buildx build --push -t $(IMAGE_NAME):amd64-linux-$$tag --platform=linux/amd64 -f $(shell if [ -e ./cmd/$*/Dockerfile.multiarch ]; then echo ./cmd/$*/Dockerfile.multiarch; else echo Dockerfile.multiarch; fi) --label revision=$(REV) .; \
                docker buildx build --push -t $(IMAGE_NAME):s390x-linux-$$tag --platform=linux/s390x -f $(shell if [ -e ./cmd/$*/Dockerfile.multiarch ]; then echo ./cmd/$*/Dockerfile.multiarch; else echo Dockerfile.multiarch; fi) --label revision=$(REV) .; \
                docker buildx build --push -t $(IMAGE_NAME):amd64-windows-$$tag --platform=windows -f $(shell if [ -e ./cmd/$*/Dockerfile.Windows ]; then echo ./cmd/$*/Dockerfile.Windows; else echo Dockerfile.Windows; fi) --label revision=$(REV) .; \
                docker manifest create --amend $(IMAGE_NAME):$$tag $(IMAGE_NAME):amd64-linux-$$tag \
                        $(IMAGE_NAME):s390x-linux-$$tag \
                        $(IMAGE_NAME):amd64-windows-$$tag; \
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

push-multiarch: $(CMDS:%=push-multiarch-%)
