
BINARY_HOME =target


all: build_ticket

clean_ticket:
	@echo "Cleaning ticket..."
	@rm -rf $(BINARY_HOME)/bin/ticket
	@mkdir -p ${BINARY_HOME}/bin

build_ticket: clean_ticket
	@echo "Building ticket..."
	@go build -o ${BINARY_HOME}/bin/ticket ./cmd/ticket/main.go


run_ticket: build_ticket
	@echo "Running ticket..."
	@${BINARY_HOME}/bin/ticket