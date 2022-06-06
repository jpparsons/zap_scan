
BINARY=sendreport

VERSION=1.0.0

LDFLAGS=-ldflags "-X sendreport/core.Version=${VERSION}"

all:
	env GOOS=linux GOARCH=386 go build ${LDFLAGS} -o ${BINARY} *.go
