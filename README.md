# zwiter
Zwiter is a very simple system with different services for `users` and `posts` using simple json files for data storage and grpc and protocol buffers for communication between the services

# Running
For running zwitter follow following steps

### 1 - Clone this repository

    go get -u github.com/dpakach/zwiter
### 2 - Change directory to the project path

    cd $GOPATH/src/github.com/dpakach/zwiter
### 3 - Initialize store files

    make initialize
### 4 - Run the Services

#### a. For User service

    make users

#### b. For Posts service
    
    make posts
        
## Running using Docker
For running using docker refer to [ this guide ](https://github.com/dpakach/zwiter/blob/master/docker_run.md)

## TODO
1. Add Authentication or Authorization
2. Refactor Store into a different Service
3. CLI client