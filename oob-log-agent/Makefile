# Go parameters


GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

AGENT=oob-log-agent
ALERT=oob-log-alert

.PHONY: all
all: build
install:
	@echo "@Install to ${GOPATH}/bin"
	mkdir -p $(GOPATH)/bin
	mkdir -p $(GOPATH)/conf
	mkdir -p $(GOPATH)/log
	cp -f $(AGENT) $(GOPATH)/bin/
	cp -f $(ALERT) $(GOPATH)/bin/
	cp -f agent-wrapper.sh $(GOPATH)/bin/
	cp -rf conf/ $(GOPATH)/

clean:
	@echo "@Clean workspace"
	$(GOCLEAN)
	rm -f $(GOPATH)/bin/$(AGENT)
	rm -f $(GOPATH)/bin/$(ALERT)

vendor: glide.lock glide.yaml
	@echo "@Download third party dependencies"
	glide up

build:  util/*.go object/*.go log-alert/main.go log-agent/main.go 
	@echo "@Building ..."
	$(GOBUILD) -v -gcflags="-N -l" -o $(AGENT) cps-oob-agent/log-agent
	$(GOBUILD) -v -gcflags="-N -l" -o $(ALERT) cps-oob-agent/log-alert
	


