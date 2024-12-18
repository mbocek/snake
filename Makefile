help: # show all commands
	@egrep -h '\s#\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?# "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

lint: # does basic linting
	@golangci-lint  --version
	golangci-lint run  --verbose --timeout 10m

build: # build executables
	go build -o tmp/snake main.go
