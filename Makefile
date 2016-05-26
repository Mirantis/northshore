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
# Args:
#   PACKAGE: Directory name to build. If not specified, "cmd/nshore" will be built.
#
# Example:
#   make
#   make build
#   make build PACKAGE=cmd/nshore

GO_CMD=go
GO_BUILD=$(GO_CMD) build -v
GO_INSTALL=$(GO_CMD) install -v
GO_CLEAN=$(GO_CMD) clean -i
GO_DEPS=$(GO_CMD) get -v

PACKAGE=cmd/nshore
TOP_PACKAGE_DIR := github.com/Mirantis/northshore

.PHONY: all build run install uninstall deps clean

#Name of final binary file
BINARY=nshore

all: run

build:
	$(info ************ Build $(BINARY) ************)
	$(GO_BUILD) -o $(BINARY) $(PACKAGE)/nshore.go

run:build
	$(info ************** Run $(BINARY) ************)
	./$(BINARY) run local

install:
	$(info ************ Install $(BINARY) **********)
	$(GO_INSTALL) $(TOP_PACKAGE_DIR)/$(PACKAGE)

uninstall:
	$(info ********** Uninstall $(BINARY) **********)
	$(GO_CLEAN) $(TOP_PACKAGE_DIR)/$(PACKAGE)

deps:
	$(info ***** Get dependencies for $(BINARY) ****)
	$(GO_DEPS) github.com/tools/godep

clean:
	$(info ************ Clean $(BINARY) ************)
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
