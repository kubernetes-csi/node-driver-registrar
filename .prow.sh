#! /bin/bash

export CSI_PROW_BUILD_PLATFORMS="linux amd64 amd64; linux ppc64le ppc64le -ppc64le; linux s390x s390x -s390x; linux arm arm -arm; linux arm64 arm64 -arm64; linux arm arm/v7 -armv7; linux riscv64 riscv64 -riscv64; windows amd64 amd64 .exe nanoserver:1809 servercore:ltsc2019; windows amd64 amd64 .exe nanoserver:ltsc2022 servercore:ltsc2022"
. release-tools/prow.sh

main
