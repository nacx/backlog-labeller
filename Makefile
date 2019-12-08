# Copyright 2019 Ignasi Barrera
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

BIN := backlog-labeller
OUT := build/$(BIN)

default: $(OUT)  ## Build the statically linked binary

$(OUT):
	CGO_ENABLED=0 GOOS=linux go build \
		-a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo \
		-o $(OUT) github.com/nacx/backlog-labeller

LINTER := bin/golangci-lint
$(LINTER):
	wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.21.0

.PHONY: lint
lint: $(LINTER)  ## Lint the project
	bin/golangci-lint run --config golangci.yml

.PHONY: clean
clean:  ## Clean all build artifacts
	rm -rf build

.PHONY: help
help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"}; \
			/^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
