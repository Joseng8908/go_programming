BINARY_NAME=micro_container
GO_FILES=./cmd/container/main.go
ROOTFS_DIR=./rootfs
CGROUP_DIR=/sys/fs/cgroup/my-container

.PHONY: all build run clean test-mem

all: build
	
build:
	@echo "빌드중: $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) $(GO_FILES)
