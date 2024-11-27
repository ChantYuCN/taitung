GO_BUILD_CMD = go
OAPI_CODEGEN_VERSION := v1.12.4
GOPATH         = $(shell go env GOPATH)
ENV_PATH       = "$(shell echo "${PATH}")":$(GOPATH)/bin:$(BINDIR)
BINDIR         = "${HOME}"/_workspace/bin
PROTOFILE      = "proto/testgrpcsvc.proto"
HOME           ?= ${HOME}
PWD            = $(shell pwd)

IMG_VERSION ?= 0.1
IMG_NAME                ?= hl-artifact-svc
DOCKER_ENV              := DOCKER_BUILDKIT=1
DOCKER_REGISTRY         ?= chant
DOCKER_REPOSITORY       ?= habana.ai
DOCKER_TAG              := ${DOCKER_REGISTRY}/${DOCKER_REPOSITORY}/${IMG_NAME}
DOCKER_LABEL_VERSION    ?= ${IMG_VERSION}



all:

openapi-gen:
	$(GO_BUILD_CMD) install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest


rest-api: openapi-gen
	${GOPATH}/bin/oapi-codegen --config=tool/oapi.yaml.server api/oapi-spec.yaml
	${GOPATH}/bin/oapi-codegen --config=tool/oapi.yaml.client api/oapi-spec.yaml

swag-spec:
	docker pull quay.io/goswagger/swagger
	docker run --rm -it  --user 1000:1000 -e GOPATH=$(GOPATH):/go -v ${HOME}:${HOME} -w $(PWD) quay.io/goswagger/swagger init spec --version 0.1.0 --format=yaml --
	
go-build:
	$(GO_BUILD_CMD) build -o bin/grestserver cmd/main.go

# rootlesskit buildkitd

nerdctl-build:
	${DOCKER_ENV} nerdctl build . -f Dockerfile.golang \
	  -t ${DOCKER_TAG}:${IMG_VERSION}
