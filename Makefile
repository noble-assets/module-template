.PHONY: proto-format proto-lint proto-gen license
all: proto-all format lint license

#=============================================================================#
#                                  Protobuf                                   #
#=============================================================================#

BUF_VERSION=1.50
BUILDER_VERSION=0.15.3

proto-all: proto-format proto-lint proto-gen

proto-format:
	@echo "==================================================================="
	@echo "Running protobuf formatter..."
	@docker run --rm --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) format --diff --write
	@echo "Completed protobuf formatting!"

proto-gen:
	@echo "==================================================================="
	@echo "Generating code from protobuf..."
	@docker run --rm --volume "$(PWD)":/workspace --workdir /workspace \
		ghcr.io/cosmos/proto-builder:$(BUILDER_VERSION) sh ./proto/generate.sh
	@echo "Completed code generation!"

proto-lint:
	@echo "==================================================================="
	@echo "Running protobuf linter..."
	@docker run --rm --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) lint
	@echo "Completed protobuf linting!"


#=============================================================================#
#                                 Tooling                                     #
#=============================================================================#

FILES := $(shell find . -name "*.go" -not -path "./simapp/*" -not -name "*.pb.go" -not -name "*.pb.gw.go" -not -name "*.pulsar.go")
license:
	@echo "==================================================================="
	@echo "Adding license to files..."
	@go-license --config .github/license.yaml $(FILES)
	@echo "Completed license addition!"

format:
	@echo "==================================================================="
	@echo "Running formatters..."
	@go tool golangci-lint fmt -c ./.golangci.yaml
	@echo "Completed formatting!"

lint:
	@echo "==================================================================="
	@echo "Running linter..."
	@go tool golangci-lint run -c ./.golangci.yaml
	@echo "Completed linting!"
