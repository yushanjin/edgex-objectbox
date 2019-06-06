#
# Copyright (c) 2018 Cavium
#
# SPDX-License-Identifier: Apache-2.0
#


.PHONY: build clean test docker run


GO=CGO_ENABLED=0 GO111MODULE=on go
GOCGO=CGO_ENABLED=1 GO111MODULE=on go

DOCKERS=docker_build_base docker_volume docker_config_seed docker_export_client docker_export_distro docker_core_data docker_core_metadata docker_core_command docker_support_logging docker_support_notifications docker_sys_mgmt_agent docker_support_scheduler
.PHONY: $(DOCKERS)

MICROSERVICES=cmd/config-seed/config-seed cmd/export-client/export-client cmd/export-distro/export-distro cmd/core-metadata/core-metadata cmd/core-data/core-data cmd/core-command/core-command cmd/support-logging/support-logging cmd/support-notifications/support-notifications cmd/sys-mgmt-executor/sys-mgmt-executor cmd/sys-mgmt-agent/sys-mgmt-agent cmd/support-scheduler/support-scheduler

.PHONY: $(MICROSERVICES)

VERSION=$(shell cat ./VERSION)
DOCKER_TAG=$(VERSION)-dev

GOFLAGS=-ldflags "-X github.com/objectbox/edgex-objectbox.Version=$(VERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

build: $(MICROSERVICES)

cmd/config-seed/config-seed:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/config-seed

cmd/core-metadata/core-metadata:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/core-metadata

cmd/core-data/core-data:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/core-data

cmd/core-command/core-command:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/core-command

cmd/export-client/export-client:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/export-client

cmd/export-distro/export-distro:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/export-distro

cmd/support-logging/support-logging:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/support-logging

cmd/support-notifications/support-notifications:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/support-notifications

cmd/sys-mgmt-executor/sys-mgmt-executor:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/sys-mgmt-executor

cmd/sys-mgmt-agent/sys-mgmt-agent:
	$(GO) build $(GOFLAGS) -o $@ ./cmd/sys-mgmt-agent

cmd/support-scheduler/support-scheduler:
	$(GOCGO) build $(GOFLAGS) -o $@ ./cmd/support-scheduler

clean:
	rm -f $(MICROSERVICES)

test:
	GO111MODULE=on go test -coverprofile=coverage.out ./...
	GO111MODULE=on go vet ./...
	gofmt -l .
	[ "`gofmt -l .`" = "" ]

run:
	cd bin && ./edgex-launch.sh

run_docker:
	bin/edgex-docker-launch.sh $(EDGEX_DB)

docker: $(DOCKERS)

docker_build_base:
	docker build \
		-f build/base/Dockerfile \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-build-base:$(GIT_SHA) \
		.

docker_volume:
	docker build \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/edge-volume:$(GIT_SHA) \
		-t objectboxio/edge-volume:$(DOCKER_TAG) \
		build/volume

docker_config_seed:
	docker build \
		-f cmd/config-seed/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/docker-core-config-seed-go:$(GIT_SHA) \
		-t objectboxio/docker-core-config-seed-go:$(DOCKER_TAG) \
		.

docker_core_metadata:
	docker build \
		-f cmd/Dockerfile --build-arg service=core-metadata \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/docker-core-metadata-go:$(GIT_SHA) \
		-t objectboxio/docker-core-metadata-go:$(DOCKER_TAG) \
		.

docker_core_data:
	docker build \
		-f cmd/Dockerfile --build-arg service=core-data \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/docker-core-data-go:$(GIT_SHA) \
		-t objectboxio/docker-core-data-go:$(DOCKER_TAG) \
		.

docker_core_command:
	docker build \
		-f cmd/core-command/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/docker-core-command-go:$(GIT_SHA) \
		-t objectboxio/docker-core-command-go:$(DOCKER_TAG) \
		.

docker_export_client:
	docker build \
		-f cmd/Dockerfile --build-arg service=export-client \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/docker-export-client-go:$(GIT_SHA) \
		-t objectboxio/docker-export-client-go:$(DOCKER_TAG) \
		.

docker_export_distro:
	docker build \
		-f cmd/export-distro/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/docker-export-distro-go:$(GIT_SHA) \
		-t objectboxio/docker-export-distro-go:$(DOCKER_TAG) \
		.

docker_support_logging:
	docker build \
		-f cmd/support-logging/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/docker-support-logging-go:$(GIT_SHA) \
		-t objectboxio/docker-support-logging-go:$(DOCKER_TAG) \
		.

docker_support_notifications:
	docker build \
		-f cmd/Dockerfile --build-arg service=support-notifications \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/docker-support-notifications-go:$(GIT_SHA) \
		-t objectboxio/docker-support-notifications-go:$(DOCKER_TAG) \
		.

docker_support_scheduler:
	docker build \
		-f cmd/Dockerfile --build-arg service=support-scheduler \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/docker-support-scheduler-go:$(GIT_SHA) \
		-t objectboxio/docker-support-scheduler-go:$(DOCKER_TAG) \
		.

docker_sys_mgmt_agent:
	docker build \
		-f cmd/sys-mgmt-agent/Dockerfile \
		--build-arg git_sha=$(GIT_SHA) \
		--label "git_sha=$(GIT_SHA)" \
		-t objectboxio/sys-mgmt-agent-go:$(GIT_SHA) \
		-t objectboxio/sys-mgmt-agent-go:$(DOCKER_TAG) \
		.
