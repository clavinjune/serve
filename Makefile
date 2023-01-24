# Copyright 2021 ClavinJune/serve
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

include tools.mk

build/out:
	@mkdir -p out/

build/darwin: build/out
	@GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 \
	go build \
		-ldflags "-s -w" \
		-o "out/darwin-amd64" main.go && \
	upx -9 out/darwin-amd64

build/docker: build/linux
	@cp out/linux-amd64 serve
	@docker build -t serve:$(shell git rev-parse HEAD | cut -c1-7) -t serve:latest .
	@rm -rf serve

build/linux: build/out
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build \
		-ldflags "-s -w" \
		-o "out/linux-amd64" main.go && \
	upx -9 out/linux-amd64

build/windows: build/out
	@GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
	go build \
		-ldflags "-s -w" \
		-o "out/windows-amd64.exe" main.go && \
	upx -9 out/windows-amd64.exe

check:
	@go run $(releaser) check
	@go run $(licenser) verify
	@go run $(linter) run
	@go run $(govulncheck) ./...

ci/release:
	@go run $(releaser) release --rm-dist

cleanup:
	@rm -rf out/ dist/
	@docker rmi -f $(shell docker images --filter "label=builder" -q)

fmt:
	@gofmt -w -s .
	@go run $(goimports) -w .
	@go vet ./...
	@go mod tidy
	@go run $(licenser) apply -r "ClavinJune/serve" 2> /dev/null

release:
	@go run $(releaser) release --rm-dist

snapshot:
	@mkdir -p src/ && go run $(releaser) release --rm-dist --snapshot && rm -rf src/
