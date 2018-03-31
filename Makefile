build:
	@echo " >> building binaries..."
	@go build -o frontendapp app/frontend/frontendapp.go
	@echo " >> frontend app builded."
	@go build -o loginapp app/login/loginapp.go
	@echo " >> loginapp app builded."
	@go build -o oauthapp app/oauth/oauthapp.go
	@echo " >> oauthapp app builded."
	@go build -o registerapp app/register/registerapp.go
	@echo " >> registerapp app builded."
	@go build -o userserviceapp app/userservice/userserviceapp.go
	@echo " >> userserviceapp app builded."

run:
	@echo "run all app..."
	@./frontendapp
	@./loginapp
	@./oauthapp
	@./registerapp
	@./userserviceapp
