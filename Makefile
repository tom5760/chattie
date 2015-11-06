# The directory this Makefile is located in
SCRIPT_DIR = $(patsubst %/,%,$(dir $(abspath $(lastword $(MAKEFILE_LIST)))))

PACKAGE_NAME = chattie

# Set GOPATH to the current directory
export GOPATH = $(SCRIPT_DIR):$(SCRIPT_DIR)/vendor

.PHONY = build clean fmt test get

build:
	go install $(PACKAGE_NAME)

clean:
	-rm -r $(SCRIPT_DIR)/bin
	-rm -r $(SCRIPT_DIR)/pkg
	-rm -r $(SCRIPT_DIR)/vendor/bin
	-rm -r $(SCRIPT_DIR)/vendor/pkg

fmt:
	go fmt $(PACKAGE_NAME)

test:
	go vet $(PACKAGE_NAME)
	go test $(PACKAGE_NAME)

get:
	export GOPATH=$(SCRIPT_DIR)/vendor; go get -u $(PKG)
	find $(SCRIPT_DIR)/vendor -name ".git" -type d | xargs rm -r
