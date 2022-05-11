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
	@bash ../build/verify_generate.sh ${OUTPUT_DIR} || \
		(echo "Error-Generated files not checked in. Run \"make gen\" and then check-in generated files." && false)


.PHONY: api_files
api_files:
	@echo "Running $@"
	rm -rf pkg/client
	@echo "Generating client from OpenAPI file..."
	openapi-generator generate --remove-operation-id-prefix -g go -o ${OUTPUT_DIR} -i ./api/swagger/swagger.yaml -p enumClassPrefix=true -p structPrefix=true -p packageName=client
	# The generator produces go.mod and go.sum, but we don't want them.
	rm -f pkg/client/go.mod pkg/client/go.sum
	# additional files that we don't need
	rm -f pkg/client/git_push.sh pkg/client/.travis.yml pkg/client/.openapi-generator-ignore


# fixup invokes a script to generate interfaces for all of the
# generated client types to allow consumers to use mocks to test their code that
# interacts with Quake via the client and adds copyrights to all the generated
# files.
.PHONY: fixup
fixup: api_files vendor
	@bash fixup_generated_files.bash ${OUTPUT_DIR}
	@bash ../build/verify_generate.sh ${GO_MOCK_DIR} || \
		(echo "Error-Generated interfaces not checked in. Run \"make fixup\" and then check-in generated files." && false)
