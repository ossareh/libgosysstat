all: build

PROJNAME := $(shell basename ${PWD})
TMPDIR := $(shell mktemp -d -t ${PROJNAME}-build-XXXX)
SRCDIR = ${TMPDIR}/src/github.com/ossareh

export GOBIN=${PWD}/libexec
export GOPATH := ${TMPDIR}:${GOPATH}

.PHONY: build
build::
	mkdir -p ${SRCDIR}
	cp -r ${PWD} ${SRCDIR}
	GOARCH= GOOS= go install -v github.com/ossareh/${PROJNAME}
