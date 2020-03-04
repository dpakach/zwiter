.PHONY: server-post server-user server-all initialize-users users posts build build-client

POSTS_FILE = data/posts.json
USERS_FILE = data/users.json
TOKENS_FILE = data/tokens.json

build:
	@go build -o ./build/posts posts/server/server.go
	@go build -o ./build/users users/server/server.go

build-client:
	@go build -o ./build/zclient cmd/zclient/main.go

posts: build
	./build/posts data/posts.json

users: build
	./build/users data/users.json

auth: build
	./build/users data/tokens.json

server-post:
	go run posts/server/server.go $(POSTS_FILE)

server-user:
	go run users/server/server.go $(USERS_FILE)

server-auth:
	go run auth/server/server.go $(TOKENS_FILE)

server-all:
	@echo "Running all servers"
	make -j 3 server-post server-user server-auth

generate-pb:
	protoc ./posts/postspb/posts.proto --go_out=plugins=grpc:.
	protoc ./users/userspb/users.proto --go_out=plugins=grpc:.
	protoc ./auth/authpb/auth.proto --go_out=plugins=grpc:.

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

initialize-auth:
	@if test -f $(TOKENS_FILE); then \
		> $(TOKENS_FILE); \
	fi; \
	echo "{\"tokens\": []}" >> $(TOKENS_FILE);

initialize: initialize-users initialize-posts initialize-auth
	@echo "Initializing Store files"

clean:
	@rm $(POSTS_FILE)
	@rm $(USERS_FILE)
	@rm $(TOKENS_FILE)
