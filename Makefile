.PHONY: server-post server-user server-all initialize-users users posts build build-client

POSTS_FILE = posts.json
USERS_FILE = users.json

build:
	@go build -o ./build/posts posts/server/server.go
	@go build -o ./build/users users/server/server.go

build-client:
	@go build -o ./build/zclient cmd/zclient/main.go

posts: build
	./build/posts posts.json

users: build
	./build/users users.json

server-post:
	go run posts/server/server.go $(POSTS_FILE)

server-user:
	go run users/server/server.go $(USERS_FILE)

server-all:
	@echo "Running all servers"
	make -j 2 server-post server-user

generate-pb:
	protoc ./posts/postspb/posts.proto --go_out=plugins=grpc:.
	protoc ./users/userspb/users.proto --go_out=plugins=grpc:.

initialize-posts:
	@if test -f $(POSTS_FILE); then \
		> $(POSTS_FILE); \
	fi; \
	echo "{\"posts\": []}" >> $(POSTS_FILE);

initialize-users:
	@if test -f $(USERS_FILE); then \
		> $(USERS_FILE); \
	fi; \
	echo "{\"users\": []}" >> $(USERS_FILE);

initialize: initialize-users initialize-posts
	@echo "Initializing Store files"

clean:
	@rm $(POSTS_FILE)
	@rm $(USERS_FILE)
