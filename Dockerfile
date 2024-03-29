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

FROM gcr.io/distroless/static:latest-arm64@sha256:6e4fa5dc7eacc9f51fac13fb4080df3e9a8d1f8d3e7e57ce881f3e10fd237620
WORKDIR /app
COPY . .
ENTRYPOINT [ "/app/serve", "-r", "/app/src" ]
