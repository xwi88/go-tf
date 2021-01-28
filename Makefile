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

default: version

all: test

clean:
	rm -r build/bin

test:
	go test -v ${BASEDIR}/test/tf_version_test.go

run: app
	${BASEDIR}/build/bin/app start
app:
	go build -v -tags ${buildTags} -ldflags ${ldFlagsDebug} -o ${BASEDIR}/build/bin/app  ${BASEDIR}
	@echo "Done app built remain gdb info"
app-darwin:
	export CGO_ENABLED=0 && export GOOS=darwin && export GOARCH=amd64 && \
	go build -v -tags ${buildTags} -ldflags ${ldFlagsRelease} -o ${BASEDIR}/build/bin/app-darwin ${BASEDIR}
	@echo "Done app built for darwin, remain gdb info "
app-linux:
	export CGO_ENABLED=0 && export GOOS=linux && export GOARCH=amd64 && \
	go build -v -tags ${buildTags} -ldflags ${ldFlagsRelease} -o ${BASEDIR}/build/bin/app-linux ${BASEDIR}
	@echo "Done app built for linux"
release:
	go build -v -tags ${buildTags} -ldflags ${ldFlagsRelease} -o ${BASEDIR}/dist/app ${BASEDIR}
	@echo "Done app release built"
release-vendor:
	go build -v -mod=vendor -tags ${buildTags} -ldflags ${ldFlagsRelease} -o ${BASEDIR}/dist/app ${BASEDIR}
	@echo "Done app release built"
version:
	${BASEDIR}/build/bin/app version
upx: app-darwin app-linux
	upx ${BASEDIR}/build/bin/app-darwin
	upx ${BASEDIR}/build/bin/app-linux
	ls -lhr ${BASEDIR}/build/bin/*
wrk:
	bash ${BASEDIR}/tool/wrk_api.sh
wrk2:
	bash ${BASEDIR}/tool/wrk_tf_predict.sh
wrk3:
	bash ${BASEDIR}/tool/wrk_tf_version.sh
wrk4:
	bash ${BASEDIR}/tool/wrk_tf_serving.sh

bench:
	go test -v -bench=. -cpu=1,2,4,8 -count=4 -benchtime=5s -benchmem -run=Benchmark_SimpleModel2 test/simple_model2_test.go | tee ${BASEDIR}/build/old.txt
stats:
	benchstat ${BASEDIR}/build/old.txt
