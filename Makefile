# Copyright 2016 The NorthShore Authors All rights reserved.
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

# Targets (see each target for more information):
#   build: Build code.
#   run: Build code and run NorthShore locally.
#   install: Install NorthShore to your $GOPATH/bin.
#   uninstall: Uninstall NorthShore from your $GOPATH/bin.
#   deps: Download and install Godep.
#   clean: Delete NorthShore binary file.

# Build code.
#
# Example:
#   make
#   make build

GO_CMD			= go
GO_BUILD		= $(GO_CMD) build -v
GO_TEST			= $(GO_CMD) test -v
GO_INSTALL		= $(GO_CMD) install -v
GO_CLEAN		= $(GO_CMD) clean
GO_DEPS			= $(GO_CMD) get -d -v
GO_FMT			= $(GO_CMD) fmt -x
GO_OS			= `uname -s | tr A-Z a-z`
PIPELINE		= examples/pipeline.yaml

PACKAGE			:= github.com/Mirantis/northshore
PKGS			= `go list ./... | grep -v /vendor/`

#Name of final binary file
BINARY			= northshore

.PHONY: all build run install uninstall deps clean

all: run

build:test
	@echo "************ Build $(BINARY) ************"
	GOOS=$(GO_OS) $(GO_BUILD) -o $(BINARY)

run:build
	@echo "************** Run $(BINARY) ************"
	./$(BINARY) run local -f $(PIPELINE)

test:
	@echo "************** Test $(BINARY) ************"
	$(GO_TEST) $(PKGS)

fmt:
	@echo "************** Format code **************"
	$(GO_FMT) $(PKGS)

install:
	@echo "************ Install $(BINARY) **********"
	$(GO_INSTALL) $(PACKAGE)

uninstall:
	@echo "********** Uninstall $(BINARY) **********"
	$(GO_CLEAN) -i $(PACKAGE)

deps:
	@echo "***** Get dependencies for $(BINARY) ****"
	$(GO_DEPS) ./...

clean:
	@echo "************ Clean $(BINARY) ************"
	$(GO_CLEAN)
