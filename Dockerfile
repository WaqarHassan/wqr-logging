# syntax=docker/dockerfile:1



FROM golang:1.19

# ENV http_proxy=http://10.123.0.132:3128
# ENV https_proxy=http://10.123.0.132:3128
# ENV GOOGLE_APPLICATION_CREDENTIALS=./dbg-corpit-dev-2c1cb73a-9c4787a8be2d.json

# Set destination for COPY
WORKDIR /app

# Download Go modules
# COPY dbg-corpit-dev-2c1cb73a-9c4787a8be2d.json ./dbg-corpit-dev-2c1cb73a-9c4787a8be2d.json
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /wqr-pdl-hello

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080


# Run
CMD ["/wqr-pdl-hello"]


