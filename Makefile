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
#   test: Run tests.
#   clean: Clean up.
#   docs: Build/Run docs (md -> html)

# Build code.
#
# Args:
#   WHAT: Directory names to build. If not specified, "everything" will be built.
#
# Example:
#   make
#   make build
#   make all WHAT=cmd/nshore
build:
	build/build.sh $(WHAT)
.PHONY: build
