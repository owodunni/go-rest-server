# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/owodunni/go-rest-server

RUN apt-get update -y; \
    apt-get install swig -y; \
    go get github.com/gorilla/mux; \
    go get github.com/owodunni/simplelib; \
    go install github.com/owodunni/go-rest-server

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/go-rest-server

# Document that the service listens on port 8080.
EXPOSE 8080
