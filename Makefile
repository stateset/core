#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags --always) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
BINDIR ?= $(GOPATH)/bin
BUILDDIR ?= $(CURDIR)/build

# Build tags
build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc cleveldb
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# Build flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=statesetd \
	-X github.com/cosmos/cosmos-sdk/version.AppName=statesetd \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

###############################################################################
###                                  Build                                 ###
###############################################################################

all: install

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/statesetd.exe ./cmd/statesetd
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/statesetd ./cmd/statesetd
endif

build-reproducible: go.sum
	$(DOCKER) buildx create --name statesetd-builder || true
	$(DOCKER) buildx use statesetd-builder
	$(DOCKER) buildx build \
		--platform linux/amd64,linux/arm64 \
		-t statesetd:local-amd64 \
		--load \
		-f Dockerfile .
	$(DOCKER) rm -f statesetd-extract || true
	$(DOCKER) create -ti --name statesetd-extract statesetd:local-amd64 bash
	$(DOCKER) cp statesetd-extract:/usr/bin/statesetd $(BUILDDIR)/statesetd-linux-amd64
	$(DOCKER) rm -f statesetd-extract

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/statesetd

###############################################################################
###                          Tools & Dependencies                          ###
###############################################################################

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify
	@go mod tidy

clean:
	rm -rf $(CURDIR)/build/

distclean: clean
	rm -rf vendor/

###############################################################################
###                              Documentation                             ###
###############################################################################

godocs:
	@echo "--> Wait a few seconds and visit http://localhost:6060/pkg/github.com/stateset/core"
	godoc -http=:6060

###############################################################################
###                           Tests & Simulation                           ###
###############################################################################

test: test-unit
test-all: test-unit test-ledger test-race test-cover

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ./...

test-race:
	@VERSION=$(VERSION) go test -mod=readonly -race -tags='ledger test_ledger_mock' ./...

test-ledger:
	@VERSION=$(VERSION) go test -mod=readonly -tags='cgo ledger test_ledger_mock' ./...

test-cover:
	@go test -mod=readonly -timeout 30m -race -coverprofile=coverage.txt -covermode=atomic -tags='ledger test_ledger_mock' ./...

benchmark:
	@go test -mod=readonly -bench=. ./...

###############################################################################
###                                Devdoc                                  ###
###############################################################################

build-docs:
	@cd docs && \
	while read p; do \
		(git checkout $${p} && npm install && VUEPRESS_BASE="/$${p}/" npm run build) ; \
		mkdir -p ~/output/$${p} ; \
		cp -r .vuepress/dist/* ~/output/$${p}/ ; \
		cp ~/output/$${p}/index.html ~/output/$${p}/404.html ; \
	done < versions ;

sync-docs:
	cd ~/output && \
	echo "role_arn = ${DEPLOYMENT_ROLE_ARN}" >> ~/.aws/config ; \
	echo "CI job = ${CIRCLE_BUILD_URL}" >> version.html ; \
	aws s3 sync . s3://${WEBSITE_BUCKET} --profile terraform --delete ; \
	aws cloudfront create-invalidation --distribution-id ${DISTRIBUTION_ID} --profile terraform --paths "/*" ;

###############################################################################
###                                Localnet                               ###
###############################################################################

build-docker-statesetdnode:
	$(MAKE) -C networks/local

# Run a 4-node testnet locally
localnet-start: build-linux localnet-stop
	@if ! [ -f build/node0/statesetd/config/genesis.json ]; then docker run --rm -v $(CURDIR)/build:/statesetd:Z statesetd/node "./statesetdnode testnet --v 4 -o . --starting-ip-address 192.168.10.2 --keyring-backend=test"; fi
	docker-compose up -d

# Stop testnet
localnet-stop:
	docker-compose down

# Clean testnet
localnet-clean:
	docker-compose down
	sudo rm -rf build/*

###############################################################################
###                           Development                                  ###
###############################################################################

# Initialize a single-node development blockchain
init:
	@echo "Initializing single-node development blockchain..."
	./build/statesetd init dev-node --chain-id stateset-dev
	./build/statesetd keys add validator --keyring-backend test
	./build/statesetd add-genesis-account $$(./build/statesetd keys show validator -a --keyring-backend test) 1000000000stake
	./build/statesetd gentx validator 1000000stake --chain-id stateset-dev --keyring-backend test
	./build/statesetd collect-gentxs

# Start development node
start-dev: build init
	@echo "Starting development node..."
	./build/statesetd start

# Quick development build and start
dev: build start-dev

# Reset blockchain data
reset:
	@echo "Resetting blockchain data..."
	./build/statesetd tendermint unsafe-reset-all
	$(MAKE) init

###############################################################################
###                               Helpers                                  ###
###############################################################################

# Download dependencies
deps:
	@echo "Downloading Go dependencies..."
	go mod download
	go mod tidy

# Format code
format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" | xargs gofmt -w -s

# Lint code
lint:
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" | xargs gofmt -d -s

.PHONY: all build-linux build build-reproducible install clean distclean test test-all test-unit test-ledger test-race test-cover benchmark build-docs sync-docs localnet-start localnet-stop localnet-clean init start-dev dev reset deps format lint