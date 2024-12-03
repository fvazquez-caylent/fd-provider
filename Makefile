# Set these before running prepare
PACK             := alph-aws
PACK_CAMEL       := alphAws
PROJECT          := github.com/AlphonsoCode/pulumi-provider-alph-aws

# These usually don't need to be changed
PACKDIR          := sdk
PROVIDER        := pulumi-resource-${PACK}
VERSION         ?= $(shell pulumictl get version)
PROVIDER_PATH   := provider
VERSION_PATH    := ${PROVIDER_PATH}.Version

GOPATH			:= $(shell go env GOPATH)

WORKING_DIR     := $(shell pwd)
EXAMPLES_DIR    := ${WORKING_DIR}/examples/yaml
TESTPARALLELISM := 4

OS    := $(shell uname)
SHELL := /bin/bash

default:: build gen_examples ensure

prepare:: prepare_dirs ensure build gen_examples
examples:: gen_examples
prepare_dirs::
	@if test -z "${PACK}"; then echo "PACK not set"; exit 1; fi
	@if test -z "${PROJECT}"; then echo "REPOSITORY not set"; exit 1; fi
	@if test ! -d "provider/cmd/pulumi-resource-xyzqw"; then "Project already prepared"; exit 1; fi # SED_SKIP

	mv "provider/cmd/pulumi-resource-xyzqw" provider/cmd/pulumi-resource-${PACK} # SED_SKIP

	if [[ "${OS}" != "Darwin" ]]; then \
		find . \( -path './.git' -o -path './sdk' \) -prune -o -not -name 'go.sum' -type f -exec sed -i '/SED_SKIP/!s,github.com/AlphonsoCode/pulumi-provider-xyzqw,${PROJECT},g' {} \; &> /dev/null; \
		find . \( -path './.git' -o -path './sdk' \) -prune -o -not -name 'go.sum' -type f -exec sed -i '/SED_SKIP/!s/[X]yzqw/${PACK_CAMEL}/g' {} \; &> /dev/null; \
		find . \( -path './.git' -o -path './sdk' \) -prune -o -not -name 'go.sum' -type f -exec sed -i '/SED_SKIP/!s/[x]yzqw/${PACK}/g' {} \; &> /dev/null; \
		sed -i '/SED_SKIP/!s,github.com/AlphonsoCode/pulumi-provider-xyzqw,${PROJECT},g' ./sdk/go.mod; \
	fi

	# In MacOS the -i parameter needs an empty string to execute in place.
	if [[ "${OS}" == "Darwin" ]]; then \
		find . \( -path './.git' -o -path './sdk' \) -prune -o -not -name 'go.sum' -type f -exec sed -i '' '/SED_SKIP/!s,github.com/AlphonsoCode/pulumi-provider-xyzqw,${PROJECT},g' {} \; &> /dev/null; \
		find . \( -path './.git' -o -path './sdk' \) -prune -o -not -name 'go.sum' -type f -exec sed -i '' '/SED_SKIP/!s/[x]yzqw/${PACK}/g' {} \; &> /dev/null; \
		find . \( -path './.git' -o -path './sdk' \) -prune -o -not -name 'go.sum' -type f -exec sed -i '' '/SED_SKIP/!s/[X]yzqw/${PACK_CAMEL}/g' {} \; &> /dev/null; \
		sed -i '' '/SED_SKIP/!s,github.com/AlphonsoCode/pulumi-provider-xyzqw,${PROJECT},g' ./sdk/go.mod; \
	fi

ensure::
	cd provider && go mod tidy
	- [[ -d sdk ]] && cd sdk && go mod tidy
	- [[ -d sdk ]] || mkdir sdk
	cd tests && go mod tidy

provider::
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

provider_debug::
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -gcflags="all=-N -l" -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" $(PROJECT)/${PROVIDER_PATH}/cmd/$(PROVIDER))

test_provider::
	cd tests && go test -short -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM} ./...

go_sdk:: $(WORKING_DIR)/bin/$(PROVIDER)
	rm -rf sdk/go
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language go

nodejs_sdk:: VERSION := $(shell pulumictl get version --language javascript)
nodejs_sdk::
	rm -rf sdk/nodejs
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language nodejs
	cd ${PACKDIR}/nodejs/ && \
		npm install && \
		npx tsc && \
		cp ../../README.md package.json package-lock.json bin/ && \
		sed -i.bak 's/$${VERSION}/$(VERSION)/g' bin/package.json && \
		rm ./bin/package.json.bak

python_sdk:: PYPI_VERSION := $(shell pulumictl get version --language python)
python_sdk::
	rm -rf sdk/python
	pulumi package gen-sdk $(WORKING_DIR)/bin/$(PROVIDER) --language python
	cp README.md ${PACKDIR}/python/
	cd ${PACKDIR}/python/ && \
		python3 setup.py clean --all 2>/dev/null && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
		rm ./bin/setup.py.bak && \
		cd ./bin && python3 setup.py build sdist

gen_examples: gen_go_example \
		gen_nodejs_example \
		gen_python_example

gen_%_example:
	rm -rf ${WORKING_DIR}/examples/$*
	pulumi convert \
		--cwd ${WORKING_DIR}/examples/yaml \
		--logtostderr \
		--generate-only \
		--non-interactive \
		--language $* \
		--out ${WORKING_DIR}/examples/$*

define pulumi_login
    export PULUMI_CONFIG_PASSPHRASE=asdfqwerty1234; \
    pulumi login --local;
endef

up::
	$(call pulumi_login) \
	cd ${EXAMPLES_DIR} && \
	pulumi stack init dev && \
	pulumi stack select dev && \
	pulumi config set name dev && \
	pulumi up -y

down::
	$(call pulumi_login) \
	cd ${EXAMPLES_DIR} && \
	pulumi stack select dev && \
	pulumi destroy -y && \
	pulumi stack rm dev -y

devcontainer::
	git submodule update --init --recursive .devcontainer
	git submodule update --remote --merge .devcontainer
	cp -f .devcontainer/devcontainer.json .devcontainer.json

.PHONY: build

build:: provider go_sdk nodejs_sdk python_sdk

# Required for the codegen action that runs in pulumi/pulumi
only_build:: build

lint::
	for DIR in "provider" "sdk" "tests" ; do \
		pushd $$DIR && golangci-lint run -c ../.golangci.yml --timeout 10m && popd ; \
	done

install:: install_provider install_nodejs_sdk
install_provider::
	cp $(WORKING_DIR)/bin/${PROVIDER} ${GOPATH}/bin

GO_TEST 	 := go test -v -count=1 -cover -timeout 2h -parallel ${TESTPARALLELISM}

test_all:: test_provider
	cd tests/sdk/nodejs && $(GO_TEST) ./...
	cd tests/sdk/python && $(GO_TEST) ./...
	cd tests/sdk/go && $(GO_TEST) ./...

install_python_sdk::
	#target intentionally blank

install_go_sdk::
	#target intentionally blank

install_nodejs_sdk::
	-npm unlink $(WORKING_DIR)/sdk/nodejs/bin
	npm link $(WORKING_DIR)/sdk/nodejs/bin
#   remove junk files created for some reason
	rm package.json package-lock.json


create_component::
	@if test -z "${SCOPE}"; then echo "SCOPE not set"; exit 1; fi
	@if test -z "${RESOURCE}"; then echo "RESOURCE not set"; exit 1; fi
	@if test -z "${RESOURCE_STRUCT}"; then echo "RESOURCE_STRUCT not set"; exit 1; fi

	@if [ -d "./provider/pkgs/providers/${SCOPE}/${RESOURCE}" ] && [ "$$(ls -A "./provider/pkgs/providers/${SCOPE}/${RESOURCE}")" ]; then echo "./provider/pkgs/providers/${SCOPE}/${RESOURCE} Exists Already. Perhaps this resource already exists?"; exit 1; fi

	@mkdir -p ./provider/pkgs/providers/${SCOPE}/${RESOURCE}
	@cp ./provider/pkgs/providers/template/component/* ./provider/pkgs/providers/${SCOPE}/${RESOURCE}

	@for file in ./provider/pkgs/providers/${SCOPE}/${RESOURCE}/*; do \
	if [ -f "$$file" ]; then \
		if [[ "${OS}" != "Darwin" ]]; then \
			sed -i 's/package component/package ${RESOURCE}/' $$file; \
			sed -i 's/\bArgs\b/${RESOURCE_STRUCT}Args/' $$file; \
			sed -i 's/\bState\b/${RESOURCE_STRUCT}State/' $$file; \
			sed -i 's/\bComponentName\b/${RESOURCE_STRUCT}/' $$file; \
		fi; \
		if [[ "${OS}" == "Darwin" ]]; then \
			sed -i '' 's/package component/package ${RESOURCE}/' $$file; \
			sed -i '' 's/[[:<:]]Args[[:>:]]/${RESOURCE_STRUCT}Args/' $$file; \
			sed -i '' 's/[[:<:]]State[[:>:]]/${RESOURCE_STRUCT}State/' $$file; \
			sed -i '' 's/[[:<:]]ComponentName[[:>:]]/${RESOURCE_STRUCT}/' $$file; \
		fi; \
	fi; \
	done
	@echo "Created Files at ./provider/pkgs/providers/${SCOPE}/${RESOURCES}"
	@echo "You may want to add this to the registry:"
	@echo ""
	@echo "import \"${PROJECT}/provider/pkgs/providers/${SCOPE}\""
	@echo ""
	@echo "ProviderRegistryEntry{"
	@echo "	PackageName:       \"${RESOURCE}\","
	@echo "	Scope:             \"${SCOPE}\","
	@echo "	Kind:              ProviderKindComponent,"
	@echo "	InferredComponent: infer.Component[*${RESOURCE}.${RESOURCE_STRUCT}, ${RESOURCE}.${RESOURCE_STRUCT}Args, *${RESOURCE}.${RESOURCE_STRUCT}State](),"
	@echo "},"

create_native::
	@if test -z "${SCOPE}"; then echo "SCOPE not set"; exit 1; fi
	@if test -z "${RESOURCE}"; then echo "RESOURCE not set"; exit 1; fi
	@if test -z "${RESOURCE_STRUCT}"; then echo "RESOURCE_STRUCT not set"; exit 1; fi

	@if [ -d "./provider/pkgs/providers/${SCOPE}/${RESOURCE}" ] && [ "$$(ls -A "./provider/pkgs/providers/${SCOPE}/${RESOURCE}")" ]; then echo "./provider/pkgs/providers/${SCOPE}/${RESOURCE} Exists Already. Perhaps this resource already exists?"; exit 1; fi

	@mkdir -p ./provider/pkgs/providers/${SCOPE}/${RESOURCE}
	@cp ./provider/pkgs/providers/template/resource/* ./provider/pkgs/providers/${SCOPE}/${RESOURCE}

	@for file in ./provider/pkgs/providers/${SCOPE}/${RESOURCE}/*; do \
	if [ -f "$$file" ]; then \
		if [[ "${OS}" != "Darwin" ]]; then \
			sed -i 's/package resource/package ${RESOURCE}/' $$file; \
			sed -i 's/\bArgs\b/${RESOURCE_STRUCT}Args/' $$file; \
			sed -i 's/\bState\b/${RESOURCE_STRUCT}State/' $$file; \
			sed -i 's/\bResourceName\b/${RESOURCE_STRUCT}/' $$file; \
		fi; \
		if [[ "${OS}" == "Darwin" ]]; then \
			sed -i '' 's/package resource/package ${RESOURCE}/' $$file; \
			sed -i '' 's/[[:<:]]Args[[:>:]]/${RESOURCE_STRUCT}Args/' $$file; \
			sed -i '' 's/[[:<:]]State[[:>:]]/${RESOURCE_STRUCT}State/' $$file; \
			sed -i '' 's/[[:<:]]ResourceName[[:>:]]/${RESOURCE_STRUCT}/' $$file; \
		fi; \
	fi; \
	done
	@echo "Created Files at ./provider/pkgs/providers/${SCOPE}/${RESOURCES}"
	@echo "You may want to add this to the registry:"
	@echo ""
	@echo "import \"${PROJECT}/provider/pkgs/providers/${SCOPE}\""
	@echo ""
	@echo "ProviderRegistryEntry{"
	@echo "	PackageName:       \"${RESOURCE}\","
	@echo "	Scope:             \"${SCOPE}\","
	@echo "	Kind:              ProviderKindResource,"
	@echo "	InferredResource: infer.Resource[${RESOURCE}.${RESOURCE_STRUCT}, ${RESOURCE}.${RESOURCE_STRUCT}Args, ${RESOURCE}.${RESOURCE_STRUCT}State](),"
	@echo "},"