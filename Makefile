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

# Make sure we pick up any local overrides.
-include .makerc

NAME := backlog-labeller
BIN := build/$(NAME)

HUB ?= docker.io/nacx
TAG ?= $(shell git rev-parse HEAD)


##@ Build targets

default: $(BIN)   ## Build the statically linked binary

$(BIN):
	CGO_ENABLED=0 GOOS=linux go build \
		-a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo \
		-o $(BIN) github.com/nacx/backlog-labeller

.PHONY: clean
clean:   ## Clean all build artifacts
	rm -rf build

##@ Code quality

LINTER := bin/golangci-lint
$(LINTER):
	wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.21.0

.PHONY: lint
lint: $(LINTER)   ## Lint the project
	bin/golangci-lint run --config golangci.yml

##@ Release targets

.PHONY: docker-build
docker-build: $(BIN)   ## Build the Docker image
	docker build -t $(HUB)/$(NAME):$(TAG) -f Dockerfile .

.PHONY: docker-push
docker-push:   ## Push the Docker image to the configured registry
	docker push $(HUB)/$(NAME):$(TAG)

##@ Others

.PHONY: help
help:   ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} \
			/^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } \
			/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)
