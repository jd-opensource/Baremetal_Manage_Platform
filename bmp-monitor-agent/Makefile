GO = go
GOCLEAN=$(GO) clean
GOBUILD = $(GO) build
GOFMT ?= gofmt "-s"
EXE ?= $(shell $(GO) env GOEXE)

AGENT = bmp-agent
PROXY = cps-agent-proxy

PACKAGES ?= $(shell go list ./... | grep -v vendor | grep -v tools)
GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./proxy/*" -not -path "./tools/*" )
PROXY_GOFILES := $(shell find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./agent/*" -not -path "./tools/*" )

BUILD_TIME := $(shell date +%Y%m%d%H%M%S)
VERSION := $(shell git describe --tags --always)
COMMITID := $(shell git rev-parse ${VERSION} | cut -c 1-32)
# CHECKSUM := $(shell md5sum conf/agent.json.prod | cut -c 1-32)
# CHECKSUM_DEBUG := $(shell md5sum conf/agent.json.dev | cut -c 1-32)

.PHONY: all vet fmt check clean
all: $(AGENT) $(PROXY)

vet:
	go vet $(PACKAGES)

fmt:
	$(GOFMT) -w $(GOFILES)

clean:
	@rm -f $(AGENT)$(EXE) $(PROXY)$(EXE)

clean-cps-agent-proxy:
	$(GOCLEAN)
	
build-cps-agent-proxy: $(PROXY_GOFILES)
	$(GOBUILD) -v -gcflags "-N -l" -o $@$(EXE) main.go

check:
	# get all go files and run go fmt on them
	@diff=$$($(GOFMT) -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;
	
$(AGENT): $(GOFILES)
	@go run tools/config-crypt.go -k $(COMMITID) -f conf/agent.json.dev -o conf/agent.json
	@go-bindata -debug -nomemcopy -nometadata -o global/global_debug.go -pkg global -tags !release conf/
	$(GOBUILD) -v -ldflags "-X cps-agent/agent.VERSION=$(VERSION) -X cps-agent/agent.BUILD_TIME=$(BUILD_TIME) -X cps-agent/agent.COMMITID=$(COMMITID) -X cps-agent/agent.CHECKSUM=$(CHECKSUM)" -gcflags "-N -l" -o $@$(EXE) 

$(PROXY): $(PROXY_GOFILES)
	$(GOBUILD) -v -gcflags "-N -l" -o $@$(EXE) cps-agent/proxy

linux: $(GOFILES)
	@mkdir -p build/linux
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o build/linux/${AGENT} main.go
	@cp -f bmpd build/linux/
	@tar czf build/bmp-agent.tgz -C build/linux ./
	@cat bmp-agent-install.sh build/bmp-agent.tgz > build/$(AGENT).bin
	# @rm -rf build/bmp-agent.tgz build/linux 
	@chmod +x build/$(AGENT).bin

windows: $(GOFILES)
	@mkdir -p build/windows/internal
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o build/windows/internal/${AGENT}.exe ./main.go
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o build/windows/${AGENT}-install.exe ./bmp-agent-install-windows.go
	@cp -f srvany.exe build/windows/internal/
	@tar czf build/bmp-agent-windows.tgz -C build/windows ./
	@rm -rf build/windows

windows_debug: $(GOFILES)
	@mkdir -p build/windows/internal
	@go run tools/config-crypt.go -k $(COMMITID) -f conf/agent.json.dev -o conf/agent.json
	@go-bindata -nomemcopy -nometadata -o global/global_release.go -pkg global -tags release conf/agent.json
	GOARCH=amd64 GOOS=windows $(GOBUILD) -v -tags release -ldflags "-X cps-agent/agent.VERSION=$(VERSION) -X cps-agent/agent.BUILD_TIME=$(BUILD_TIME) -X cps-agent/agent.COMMITID=$(COMMITID) -X cps-agent/agent.CHECKSUM=$(CHECKSUM_DEBUG) -w -s " -o build/windows/internal/${AGENT}.exe main.go
	GOARCH=amd64 GOOS=windows $(GOBUILD) -o build/windows/${AGENT}_install.exe bmp-agent-install-windows.go
	@cp -f srvany.exe build/windows/internal/
	@tar czf build/bmp-agent-windows.tgz -C build/windows ./
	@rm -rf build/windows
	




release: linux
	
