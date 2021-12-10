GO_EXECUTABLE ?= go
HEAD = `git rev-parse HEAD`
TIME = `date +%FT%T%z`

NAME = av_converter
BINARY = ${NAME}

LDFLAGS = -ldflags "-X main.Version=${HEAD} -X main.BuildTime=${TIME}"

UNAME = $(shell uname)
ifeq (${UNAME}, Darwin)
	os=darwin
else
	os=windows
endif

ifeq (${os}, windows)
    BINARY = ${NAME}.exe
endif

build:
	${GO_EXECUTABLE} build  ${LDFLAGS} -o ${BINARY}

debug:
	${GO_EXECUTABLE} build -gcflags "all= -N  -l" ${LDFLAGS} -o ${BINARY}

clean:
	rm -f ${BINARY}

.PHONY: build debug clean
