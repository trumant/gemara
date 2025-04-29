BUILD_FLAGS=-X 'main.GitCommitHash=`git rev-parse --short HEAD`' -X 'main.BuiltAt=`date +%FT%T%z`' -X 'main.Version=`git describe --tags`'
BUILD_WIN=@env GOOS=windows GOARCH=amd64 go build -o privateer-windows.exe
BUILD_LINUX=@env GOOS=linux GOARCH=amd64 go build -o privateer-linux
BUILD_MAC=@env GOOS=darwin GOARCH=amd64 go build -o privateer-darwin

binary: tidy test build
quick: build
testcov: test test-cov
release: tidy test release-nix release-win release-mac

define build
echo "  >  Building binary ..."; \
cd $(1); \
go build -o ../../$(1) -ldflags="$(BUILD_FLAGS)";
endef

define test
echo "  >  Validating code ..."
cd $(1); \
go vet ./...
go test ./...
endef

build:
	@$(call build,layer2)
	@$(call build,layer4)

test:
	@$(call test,layer2)
	@$(call test,layer4)

tidy:
	@echo "  >  Tidying go.mod ..."
	@go mod tidy

test-cov:
	@echo "Running tests and generating coverage output ..."
	@go test ./... -coverprofile coverage.out -covermode count
	@sleep 2 # Sleeping to allow for coverage.out file to get generated
	@echo "Current test coverage : $(shell go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+') %"

release-candidate: tidy test
	@echo "  >  Building release candidate for Linux..."
	$(BUILD_LINUX) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=nix-rc'"
	@echo "  >  Building release candidate for Windows..."
	$(BUILD_WIN) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=win-rc'"
	@echo "  >  Building release for Darwin..."
	$(BUILD_MAC) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=darwin'"

release-nix:
	@echo "  >  Building release for Linux..."
	$(BUILD_LINUX) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=linux'"

release-win:
	@echo "  >  Building release for Windows..."
	$(BUILD_WIN) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=windows'"

release-mac:
	@echo "  >  Building release for Darwin..."
	$(BUILD_MAC) -ldflags="$(BUILD_FLAGS) -X 'main.VersionPostfix=darwin'"

todo:
	@read -p "Write your todo here: " TODO; \
	echo "- [ ] $$TODO" >> TODO.md

lintcue:
	@echo "  >  Linting CUE files ..."
	@echo "  >  Linting layer-2.cue ..."
	@cue eval ./schemas/layer-2.cue --all-errors --verbose
	@echo "  >  Linting layer-4.cue ..."
	@cue eval ./schemas/layer-4.cue --all-errors --verbose

lintexamples:
	@echo "  >  Linting example files ..."
	@echo "  >  Linting schemas/example_evaluation_results ..."
	@cue vet schemas/layer-4.cue schemas/example_evaluation_results.yml -d '#Layer4'

cuegen:
	@echo "  >  Generating types from cue schema ..."
	@echo "  >  Generating types for layer2 ..."
	@cue exp gengotypes ./schemas/layer-2.cue
	@mv cue_types_gen.go pkg/layer2/generated_types.go
	@echo "  >  Generating types for layer4 ..."
	@cue exp gengotypes ./schemas/layer-4.cue
	@mv cue_types_gen.go pkg/layer4/generated_types.go

PHONY: lintcue lintexamples
