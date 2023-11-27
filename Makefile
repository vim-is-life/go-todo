##
# go-todo
#
# @file
# @version 0.1

TARGET=go-todo
PORT_TO_RUN_ON=:9000
# PROGRAM_FILES="$(TARGET) run.sh README.md views/todo.gohtml"
PROGRAM_FILES=./$(TARGET) ./Makefile ./README.md ./views ./run.sh
# PROGRAM_FILES=./$(TARGET)
FAR_LOCATION=${REMOTE_LOCATION}

.PHONY: all run run-debug deploy $(TARGET)
all: $(TARGET)

$(TARGET): ./**/*.go
	etags -Q --declarations ./**/*.go
	ctags ./**/*.go
	go build

run: $(TARGET)
    # APP_PORT=$(PORT_TO_RUN_ON) ./go-todo
	APP_PORT=$(PORT_TO_RUN_ON) ./run.sh

run-debug: $(TARGET)
    # GOTRACEBACK=crash APP_PORT=$(PORT_TO_RUN_ON) ./go-todo
	APP_PORT=$(PORT_TO_RUN_ON) ./run.sh debug

# scp -r "$PROGRAM_FILES" "$REMOTE_LOCATION"
deploy:
ifeq (${FAR_LOCATION},)
	@echo "ERROR: need a ssh host and location to deploy to" >&1
	@echo "usage: REMOTE_LOCATION=<ssh-host:path/to/remote/location> make deploy"
else
	rsync -a $(PROGRAM_FILES) "$(REMOTE_LOCATION)"
endif

# end
