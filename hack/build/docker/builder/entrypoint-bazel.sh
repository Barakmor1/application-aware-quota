#!/usr/bin/env bash

#Copyright 2025 The AAQ Authors.
#
#Licensed under the Apache License, Version 2.0 (the "License");
#you may not use this file except in compliance with the License.
#You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
#Unless required by applicable law or agreed to in writing, software
#distributed under the License is distributed on an "AS IS" BASIS,
#WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#See the License for the specific language governing permissions and
#limitations under the License.

set -eo pipefail

source /etc/profile.d/gimme.sh

export JAVA_HOME=/usr/lib/jvm/java-11
export PATH=${GOPATH}/bin:/go/bin:/opt/gradle/gradle-6.6/bin:$PATH

eval "$@"
