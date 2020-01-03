# from base image golang:latest
FROM golang:latest 

# service type
ENV service users

# Configure GOPATH and project path
RUN mkdir -p $GOPATH/src/github.com/dpakach/zwiter
ADD . $GOPATH/src/github.com/dpakach/zwiter
WORKDIR $GOPATH/src/github.com/dpakach/zwiter

# Install required dependencies
RUN go get github.com/golang/protobuf/proto
RUN go get google.golang.org/grpc

# Setup project
RUN make initialize

# Expose required Ports
EXPOSE 8001
EXPOSE 8002

# Running the project
CMD make $service
