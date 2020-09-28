# Copyright 2020, Author: Ashutosh Kumar (GithubID: @sonasingh46).
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.


# This Dockerfile builds a recent metrics-server using the latest binary from
# metrics-server  releases.

FROM alpine:3.11.5

RUN apk add --no-cache \
    iproute2 \
    bash \
    curl \
    net-tools \
    mii-tool \
    procps \
    libc6-compat \
    ca-certificates

COPY metrics-server /usr/local/bin/metrics-server

ARG ARCH
ARG DBUILD_DATE
ARG DBUILD_REPO_URL

LABEL org.label-schema.name="metrics-server"
LABEL org.label-schema.description="Metrics ingestion server"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.build-date=$DBUILD_DATE
LABEL org.label-schema.vcs-url=$DBUILD_REPO_URL

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/metrics-server"]

