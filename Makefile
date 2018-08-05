usage: FORCE
	# See targets in Makefile (e.g. "buildlet.darwin-amd64")
	exit 1

FORCE:

.PHONY: build
build:
	@echo " >> building binaries..."
	@go build -o frontendapp cmd/frontend/frontendapp.go
	@echo " >> frontend app builded."
	@go build -o loginapp cmd/login/loginapp.go
	@echo " >> loginapp app builded."
	@go build -o oauthapp cmd/oauth/oauthapp.go
	@echo " >> oauthapp app builded."
	@go build -o registerapp cmd/register/registerapp.go
	@echo " >> registerapp app builded."
	@go build -o userserviceapp cmd/userservice/userserviceapp.go
	@echo " >> userserviceapp app builded."

run:
	@echo "run all app..."
	@./frontendapp
	@./loginapp
	@./oauthapp
	@./registerapp
	@./userserviceapp

.PHONY: default
default: build

.DEFAULT_GOAL := build