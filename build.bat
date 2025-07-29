go build -ldflags "-s -w"

set GOARCH=amd64
set GOOS=linux

go build -ldflags "-s -w"