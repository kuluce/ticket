
BINARY_HOME =target
OS_NAME =$(shell go env GOHOSTOS)
CC=gcc

all: build_ticket

clean_windows:
	@echo "Cleaning windows ticket..."
	@rmdir /s /q $(BINARY_HOME)
	@mkdir  ${BINARY_HOME}
	@mkdir  ${BINARY_HOME}\bin

clean_macos:
	@echo "Cleaning macos ticket..."
	@rm -rf $(BINARY_HOME)/bin/ticket
	@mkdir -p ${BINARY_HOME}/bin

clean_ticket:  
ifeq ($(OS),Windows_NT)
	@make clean_windows
else
	@make clean_macos
endif 

build_ticket: clean_ticket
	@echo "Building ticket..."
ifeq ($(OS),Windows_NT)
	@go build -o ${BINARY_HOME}\bin\ticket.exe ./cmd/ticket/main.go
else
	@go build -o ${BINARY_HOME}/bin/ticket ./cmd/ticket/main.go
endif

run_ticket: build_ticket
	@echo "Running ticket..."
	@${BINARY_HOME}/bin/ticket