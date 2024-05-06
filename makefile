BINARY_NAME=iot-wifi-setup_linux_arm64
PROJECT_DIR := $(dir $(realpath -s $(lastword $(MAKEFILE_LIST))))

build:
	GOOS=linux GOARCH=arm64 go build -o ${BINARY_NAME}

#run:
#	go build -o ${BINARY_NAME} main.go
#	./${BINARY_NAME}

mocks:
	find "${PROJECT_DIR}" -type f -name "mock_*.go" -delete
	docker run --rm -v "${PROJECT_DIR}":/src -v "${GOPATH}/pkg/mod":/go/pkg/mod -w /src vektra/mockery:v2.43.0

clean:
	go clean
	rm ${BINARY_NAME}