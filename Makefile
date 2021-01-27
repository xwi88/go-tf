# Makefile to build the command lines and tests in this project.
# This Makefile doesn't consider Windows Environment. If you use it in Windows, please be careful.
SHELL := /bin/sh

existBash = $(shell cat /etc/shells|grep -w /bin/bash|grep -v grep)
ifneq (, $(strip ${existBash}))
	SHELL = /bin/bash
endif
$(info shell will use ${SHELL})

#BASEDIR = $(shell pwd)
BASEDIR = $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

# add following lines before go build!
versionDir = github.com/xwi88/version

gitBranch = $(shell git symbolic-ref --short -q HEAD)

ifeq ($(gitBranch),)
gitTag = $(shell git describe --always --tags --abbrev=0)
endif

buildTime = $(shell date "+%FT%T%z")
gitCommit = $(shell git rev-parse HEAD)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

# -ldflags flags accept a space-separated list of arguments to pass to an underlying tool during the build.
ldFlagsDebug="-X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} \
 -X ${versionDir}.buildTime=${buildTime} -X ${versionDir}.gitCommit=${gitCommit} \
 -X ${versionDir}.gitTreeState=${gitTreeState}"

# -s -w
#ldFlagsRelease="-s -w -X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} \
#  -X ${versionDir}.buildTime=${buildTime} -X ${versionDir}.gitCommit=${gitCommit} \
#  -X ${versionDir}.gitTreeState=${gitTreeState}"

# -s -w
# -a #force rebuilding of packages that are already up-to-date.
ldFlagsRelease="-installsuffix -s -w -X ${versionDir}.gitBranch=${gitBranch} -X ${versionDir}.gitTag=${gitTag} \
  -X ${versionDir}.buildTime=${buildTime} -X ${versionDir}.gitCommit=${gitCommit} \
  -X ${versionDir}.gitTreeState=${gitTreeState}"

$(shell mkdir -p ${BASEDIR}/build/bin/conf)

#buildTags=""
buildTags="jsoniter"

.PHONY: default test

default: test

all: test

clean:
	rm -r build/bin

test:
	go build -v -tags ${buildTags} -ldflags ${ldFlagsDebug} -o ${BASEDIR}/build/bin/test ${BASEDIR}/test/
	@echo "Done test built remain gdb info"

run:
	${BASEDIR}/build/bin/test

re-run: test run
