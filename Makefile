before.build:
	go mod download && go mod vendor

build.smtrackerp:
	@echo "build in ${PWD}";go build smtrackerp.go