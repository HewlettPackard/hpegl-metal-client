# (C) Copyright 2020-2022 Hewlett Packard Enterprise Development LP

PWD := $(shell pwd)
GOPATH := $(shell go env GOPATH)

export GOPRIVATE=github.com/HewlettPackard

OUTPUT_DIR="pkg/client"

BIN_DIR := ${PWD}/bin
export IFACEMAKER := ${BIN_DIR}/ifacemaker

# the location of the go mock generator binary
export GO_MOCKGEN := ${BIN_DIR}/mockgen

# The package containing the mocks
export MOCK_CLIENT_PACKAGE=mockquakeclient

# The location of any generated golang mocks
export GO_MOCK_DIR := ${PWD}/pkg/${MOCK_CLIENT_PACKAGE}

default: build

vendor: ../go.mod ../go.sum
	go mod download
	go mod vendor

.PHONY: validate
validate: 
	@echo "Running $@"
	openapi-generator validate -i ./api/swagger/swagger.yaml 

.PHONY: gen
gen: api_files fixup

.PHONY: verify_gen
verify_gen: gen 
	@echo "Checking the status of the generated client directory"
	@bash ../build/verify_generate.sh ${OUTPUT_DIR}

.PHONY: api_files
api_files:
	@echo "Running $@"
	rm -rf pkg/client
	@echo "Generating client from swagger..."
	openapi-generator generate --remove-operation-id-prefix -g go -o ${OUTPUT_DIR} -i ./api/swagger/swagger.yaml -p enumClassPrefix=true -p structPrefix=true -p packageName=client
	# The generator produces go.mod and go.sum, but we don't want them.
	rm -f pkg/client/go.mod pkg/client/go.sum
	# additional files that we don't need
	rm -f pkg/client/git_push.sh pkg/client/.travis.yml pkg/client/.openapi-generator-ignore

${IFACEMAKER}:
	GOBIN="$(@D)" go get github.com/vburenin/ifacemaker@v1.1.0

${GO_MOCKGEN}:
	GOBIN="$(@D)" go get github.com/golang/mock/mockgen@v1.5.0

# fixup invokes a script to generate interfaces for all of the
# generated client types to allow consumers to use mocks to test their code that
# interacts with Quake via the client and adds copyrights to all the generated
# files.
.PHONY: fixup
fixup: api_files ${IFACEMAKER} ${GO_MOCKGEN} vendor
	@bash fixup_generated_files.bash ${OUTPUT_DIR}

