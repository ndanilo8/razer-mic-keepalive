FROM golang:1.23

WORKDIR /build

# Copy your source code
COPY . .

# Install dependencies
RUN go mod tidy

# Cross-compile to Windows EXE with GUI subsystem
RUN GOOS=windows GOARCH=amd64 \
    go build -ldflags="-H=windowsgui" -o RazerMicKeepAlive.exe main.go
