# Build variables
BINARY_NAME=iconhash
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
DATE=$(shell date +%FT%T%z)
BUILD_BY=$(shell whoami)

# Docker variables
DOCKER_IMAGE=cyberspacesec/iconhash
DOCKER_TAG=$(VERSION)

# Go build flags
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.date=${DATE} -X main.builtBy=${BUILD_BY}"

.PHONY: all build clean install uninstall test docker-build docker-push

all: clean build

build:
	@echo "Building ${BINARY_NAME}..."
	go build ${LDFLAGS} -o ${BINARY_NAME}

clean:
	@echo "Cleaning up..."
	rm -f ${BINARY_NAME}

install: build
	@echo "Installing ${BINARY_NAME}..."
	mv ${BINARY_NAME} ${GOPATH}/bin/

uninstall:
	@echo "Uninstalling ${BINARY_NAME}..."
	rm -f ${GOPATH}/bin/${BINARY_NAME}

test:
	@echo "Running tests..."
	go test -v ./...

# Create a sample favicon.ico for testing
sample:
	@echo "Creating sample favicon.ico for testing..."
	mkdir -p test
	curl -s -o test/favicon.ico https://www.baidu.com/favicon.ico

# Run the tool with test favicon
test-sample: build sample
	@echo "Testing with sample favicon.ico..."
	./${BINARY_NAME} -f test/favicon.ico

# Run the tool with a URL
test-url: build
	@echo "Testing with URL..."
	./${BINARY_NAME} -u https://www.baidu.com/favicon.ico

# Docker targets
docker-build:
	@echo "Building Docker image ${DOCKER_IMAGE}:${DOCKER_TAG}..."
	docker build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .
	docker tag ${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_IMAGE}:latest

docker-push: docker-build
	@echo "Pushing Docker image ${DOCKER_IMAGE}:${DOCKER_TAG}..."
	docker push ${DOCKER_IMAGE}:${DOCKER_TAG}
	docker push ${DOCKER_IMAGE}:latest

# Run the tool inside a Docker container
docker-run:
	@echo "Running in Docker container..."
	docker run --rm ${DOCKER_IMAGE}:latest $(ARGS) 