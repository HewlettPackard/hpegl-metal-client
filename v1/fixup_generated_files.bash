#! /bin/bash
COPYRIGHT="// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP"
set -e

if [ $# -ne 1 ]; then
  echo "usage: $0 <directory>" 1>&2
  exit 1
fi

OUTPUT_DIR=$1
shift

mkdir -p ${GO_MOCK_DIR}

cd ${OUTPUT_DIR}

#
# Generate interfaces for the generated api files
#
echo "Generating interfaces for the client..."
for file in $(ls api_*.go); do
  INTERFACE_FILE=$(echo ${file} | sed -e 's/^api_/interface_/g')
  MOCK_FILE=$(echo ${file} | sed -e 's/^api_/mock_/g')
  TYPE=$(grep "type \([^ ]*\) service" ${file} | sed -e 's/type \([^ ]*\)ApiService service/\1/g')
  SERVICE="${TYPE}ApiService"
  INTERFACE="${TYPE}API"
  # the generator gets the capitalization of IPPools wrong.
  if [ "${INTERFACE}" == "IppoolsAPI" ]; then
    INTERFACE="IPPoolsAPI"
  fi

  IFACEMAKER='go run github.com/vburenin/ifacemaker'
  GO_MOCKGEN='go run github.com/golang/mock/mockgen'
  
  echo "Generating interface and mocks for ${SERVICE}"
   ${IFACEMAKER} -p client -y "${INTERFACE} defines the client functions provided for ${TYPE}." --struct=${SERVICE} --file=${file} --output=${INTERFACE_FILE} --iface=${INTERFACE}
   sed -i "s/\(${TYPE}Api\) \*${SERVICE}/\1 ${INTERFACE}/g" client.go
   ${GO_MOCKGEN} --source ${INTERFACE_FILE} --destination ${GO_MOCK_DIR}/${MOCK_FILE} --package=${MOCK_CLIENT_PACKAGE} ${INTERFACE}
done

#
# Create error.go
#
echo "Generating generic-error.go..."

cat - <<*EOF* > generic-error.go
package client

// NewGenericOpenAPIError creates a GenericOpenAPIError from the provided
// parameters.
func NewGenericOpenAPIError(body []byte, error string, model interface{}) GenericOpenAPIError {
	return GenericOpenAPIError{
		body:  body,
		error: error,
		model: model,
	}
}

// Message returns the short error message for display purpose
// If e.model is set and is of type ErrorResponse, then returns the value in ErrorResponse.Message,
// otherwise, returns the entire response body in string format.  
func (e GenericOpenAPIError) Message() string {
	msg := string(e.Body())

	model := e.Model()
	if model != nil {
		res, ok := model.(ErrorResponse)
		if ok {
			msg = res.Message
		}
	}
	return msg
}
*EOF*

#
# Add copyrights to all the generated files
#
echo "Adding copyright to generated files..."
for file in $(ls *go) $(ls ${GO_MOCK_DIR}/*.go); do
  TMP_FILE="$(mktemp /tmp/copyright.XXXXXX)"
  cat - ${file} << *EOF* > ${TMP_FILE}
${COPYRIGHT}

*EOF*
  mv ${TMP_FILE} ${file}
done
