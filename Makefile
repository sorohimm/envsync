# suppress output, run `make XXX V=` to be verbose
V := @

RELEASE=$(shell git describe --always --tags)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

OUT_DIR := ./builds

ENVSYNC_APP := "envsync"

.PHONY: build
build: ENVSYNC_OUT := $(OUT_DIR)/$(ENVSYNC_APP)
build: ENVSYNC_MAIN := ./cmd
build:
	@echo BUILDING $(ENVSYNC_OUT)
	$(V)go build -ldflags "-s -w -X main.version=${RELEASE} -X main.buildTime=${BUILD_TIME}" -o $(ENVSYNC_OUT) $(ENVSYNC_MAIN)
	@echo DONE