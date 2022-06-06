
# make all

BINARY=sendreport

VERSION=1.0.0

LDFLAGS=-ldflags "-X sendreport/core.Version=${VERSION}"

all:
	env GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o ${BINARY} ${BINARY}.go
