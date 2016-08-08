DOCKER_CMD ?= docker
PREFIX    ?= /usr/bin

BIN_DIR  := $(CURDIR)/bin
BUILD_IMAGE := numa

DOCKER_VERS      := $(shell $(DOCKER_CMD) version -f '{{.Client.Version}}')
DOCKER_VERS_MAJ  := $(shell echo $(DOCKER_VERS) | cut -d. -f1)
DOCKER_VERS_MIN  := $(shell echo $(DOCKER_VERS) | cut -d. -f2)
DOCKER_SUPPORTED := $(shell [ $(DOCKER_VERS_MAJ) -eq 1 -a $(DOCKER_VERS_MIN) -ge 9 ] && echo true)

.PHONY: all build local

all: build

build:
ifneq ($(DOCKER_SUPPORTED),true)
	$(error Unsupported Docker version)
endif
	@$(DOCKER_CMD) build -t $(BUILD_IMAGE) -f Dockerfile.build $(CURDIR)
	@mkdir -p $(BIN_DIR)
	@$(DOCKER_CMD) run --rm -it --net=host -v $(BIN_DIR):/go/bin $(BUILD_IMAGE)