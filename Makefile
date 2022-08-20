BINARY_NAME=genxoft-server
WEBAPP_DIST=./web
BUILD_COUNTER_FILE=build-counter.txt
BUILD_MODE=$(mode)
ifeq ($(mode),)
	BUILD_MODE="prod"
endif

VERSION=$$(git describe --tags --abbrev=0)
RELEASE_ID=`cat ${BUILD_COUNTER_FILE}`
LDFLAGS=-ldflags "-w -s -X main.Version=${VERSION} -X main.ReleaseId=${RELEASE_ID} -X main.Mode=${BUILD_MODE}"

.PHONY: build-inc
build-inc:
	./build/incbuild.sh ${BUILD_COUNTER_FILE}
	@echo "${VERSION} build: ${RELEASE_ID}"

.PHONY: dep-client
dep-client:
	@echo "> Downloading client dependencies..."
	@npm --prefix ./frontend install

.PHONY: dep-backend
dep-backend:
	@echo "> Downloading backend dependencies..."
	@go mod download

.PHONY: build-client
build-client:
	@echo "> Building the client ..."
	REACT_APP_VERSION=${VERSION} REACT_APP_BUILD=${RELEASE_ID} BUILD_PATH=./../${WEBAPP_DIST} npm --prefix ./frontend run build

.PHONY: build-backend
build-backend:
	@echo "> Building the backend ..."
	@go build -o ${BINARY_NAME} ${LDFLAGS} ./cmd/webserver/main.go

.PHONY: create_artifact
create_artifact:
	./build/create_artifact.sh ${BUILD_MODE} ${VERSION}

.PHONY: run
run:
	./${BINARY_NAME}

.PHONY: build
build: build-inc dep-client build-client dep-backend build-backend create_artifact

.PHONY: clean
clean:
	go clean

.PHONY: test
test:
	go test -v -race ./...

.PHONY: test_coverage
test_coverage:
	go test -v -race ./... -coverprofile=coverage.out



.DEFAULT_GOAL := build