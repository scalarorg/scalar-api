TOOLS_DIR := tools

BUILDDIR ?= $(CURDIR)/build

ldflags := $(LDFLAGS)
build_tags := $(BUILD_TAGS)
build_args := $(BUILD_ARGS)

ifeq ($(VERBOSE),true)
	build_args += -v
endif

ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static" -v
endif

BUILD_TARGETS := build install
BUILD_FLAGS := --tags "$(build_tags)" --ldflags '$(ldflags)'

all: build install

build: BUILD_ARGS := $(build_args) -o $(BUILDDIR)

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

.PHONY: build install tests dev run-local run-unprocessed-events-replay-local generate-mock-interface build-swagger

build-docker:
	$(MAKE) BBN_PRIV_DEPLOY_KEY=${BBN_PRIV_DEPLOY_KEY} -C contrib/images xchains-api

start-xchains-api: build-docker stop-service
	docker-compose up -d

stop-service:
	docker-compose down
	
run-local:
	@go run cmd/xchains-api/main.go \
		--config config/config-local.yml \
		--params config/global-params.json \
		--finality-providers config/finality-providers.json
dev:
	@go run cmd/xchains-api/main.go \
		--config config/config-local.yml \
		--params config/global-params.json \
		--finality-providers config/finality-providers.json

# We don't use config, params and finality providers, it's here due to dependency reason
run-unprocessed-events-replay-local:
	./bin/local-startup.sh;
	sleep 5;
	go run cmd/xchains-api/main.go \
		--config config/config-local.yml \
		--params config/global-params.json \
		--finality-providers config/finality-providers.json \
		--replay

generate-mock-interface:
	cd internal/db && mockery --name=DBClient --output=../../tests/mocks --outpkg=dbmock --filename=mock_db_client.go

tests:
	./bin/local-startup.sh;
	go test -v -cover -p 1 ./... -count=1


build-swagger:
	swag init --parseDependency --parseInternal -d cmd/xchains-api,internal/api,internal/types