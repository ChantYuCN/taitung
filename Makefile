GO_BUILD_CMD = go
PROTOC_GEN_VER = v0.10.1
OAPI_CODEGEN_VERSION := v1.12.4
GOPATH         = $(shell go env GOPATH)
ENV_PATH       = "$(shell echo "${PATH}")":$(GOPATH)/bin:$(BINDIR)
BINDIR         = "bin"
PROTOFILE      = "proto/mo.proto"
HOME           ?= ${HOME}
PWD            = $(shell pwd)

IMG_VERSION ?= 0.2
IMG_NAME                ?= hl-artifact-svc
DOCKER_ENV              := DOCKER_BUILDKIT=1
DOCKER_REGISTRY         ?= chant
DOCKER_REPOSITORY       ?= habana.ai
DOCKER_TAG              := ${DOCKER_REGISTRY}/${DOCKER_REPOSITORY}/${IMG_NAME}
DOCKER_LABEL_VERSION    ?= ${IMG_VERSION}



all:

prep:
	@if [ ! -f ${BINDIR}/protoc ]; then \
    rm -f protoc-21.5-linux-x86_64.zip* \
    && wget https://github.com/protocolbuffers/protobuf/releases/download/v21.5/protoc-21.5-linux-x86_64.zip \
    && mkdir -p ${BINDIR}/ \
    && unzip -o protoc-21.5-linux-x86_64.zip -d ${BINDIR}/ \
    && rm -f protoc-21.5-linux-x86_64.zip*; \
    fi

proto-gen: prep
	$(GO_BUILD_CMD) install github.com/google/gnostic@latest
	$(GO_BUILD_CMD) install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	$(GO_BUILD_CMD) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	$(GO_BUILD_CMD) install github.com/googleapis/gnostic-grpc@latest
	$(GO_BUILD_CMD) install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
	$(GO_BUILD_CMD) install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	$(GO_BUILD_CMD) install github.com/envoyproxy/protoc-gen-validate@${PROTOC_GEN_VER}

pbstub-generate: proto-gen
	cd api && PATH=${ENV_PATH} ../${BINDIR}/bin/protoc -I proto/.   \
            --go_out=./proto --go_opt=paths=source_relative --go-grpc_out=./proto --go-grpc_opt=paths=source_relative \
            --grpc-gateway_out=allow_delete_body=true:. --grpc-gateway_opt=paths=source_relative \
            ${PROTOFILE}
	cd -


openapi-gen:
	$(GO_BUILD_CMD) install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest


rest-api: openapi-gen
	${GOPATH}/bin/oapi-codegen --config=tool/oapi.yaml.server api/oapi-spec.yaml
	${GOPATH}/bin/oapi-codegen --config=tool/oapi.yaml.client api/oapi-spec.yaml

swag-spec:
	docker pull quay.io/goswagger/swagger
	docker run --rm -it  --user 1000:1000 -e GOPATH=$(GOPATH):/go -v ${HOME}:${HOME} -w $(PWD) quay.io/goswagger/swagger init spec --version 0.1.0 --format=yaml --
	
go-build:
	$(GO_BUILD_CMD) build -o ${BINDIR}/grestserver cmd/main.go
	$(GO_BUILD_CMD) build -o ${BINDIR}/grpcclient client/grpcclient.go

# rootlesskit buildkitd

nerdctl-build:
	${DOCKER_ENV} nerdctl build . -f Dockerfile.golang \
	  -t ${DOCKER_TAG}:${IMG_VERSION}
