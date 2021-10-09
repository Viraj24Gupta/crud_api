.PHONY: all
all: compile start

APP=crud_api
APP_EXECUTABLE="./bin/$(APP)"

compile:
	mkdir -p bin/
	go build -o $(APP_EXECUTABLE)
	@echo "binary created in bin directory"

start:
	@echo "server started"
	bin/$(APP)
