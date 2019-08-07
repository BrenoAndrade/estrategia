init-install: 
	@make deps
	@make redis
	@make init

init: 
	go run main.go

deps:
	go get github.com/gorilla/handlers
	go get github.com/gorilla/mux
	go get github.com/go-redis/redis
	go get github.com/watson-developer-cloud/go-sdk/naturallanguageunderstandingv1

redis: 
	sudo docker run -d -p 6379:6379 -i -t redis:3.2.5-alpine