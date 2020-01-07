# from base image golang:latest
FROM golang:latest

# service type
ENV service users
ENV USER_HOST zusers
ENV USER_PORT 8002
ENV POST_HOST zposts
ENV POST_PORT 8001


# Configure GOPATH and project path
RUN mkdir -p $GOPATH/src/github.com/dpakach/zwiter
ADD . $GOPATH/src/github.com/dpakach/zwiter
WORKDIR $GOPATH/src/github.com/dpakach/zwiter

# Install required dependencies
RUN go get github.com/golang/protobuf/proto
RUN go get google.golang.org/grpc

# Setup project
RUN make initialize

# Running the project
CMD make $service
