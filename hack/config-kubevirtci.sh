#
# This file is part of the KubeVirt project
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
#
# Copyright the KubeVirt Authors.
#
#

BASE_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
KUBEVIRTCI_PATH="${BASE_DIR}/kubevirtci/cluster-up/"
KUBEVIRTCI_CONFIG_PATH="${BASE_DIR}/kubevirtci/_ci-configs"
KUBEVIRT_RELEASE=${KUBEVIRT_RELEASE:-"latest_nightly"}