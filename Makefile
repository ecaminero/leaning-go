PID_FILE = /tmp/my-app.pid
GO_FILES = $(wildcard *.go)

# Start task performs "go run main.go" command and writes it's process id to PID_FILE.
start:
	go run ./src/$(GO_FILES) & echo $$! > $(PID_FILE)

# Serve task will run fswatch monitor and performs restart task if any source file changed. Before serving it will execute start task.
serve: start
	fswatch -or --event=Updated ./src | \
	xargs -n1 -I {}
  
.PHONY: serve