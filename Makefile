# All bootstrap starts with uname
UNAME_OS=$(shell uname -s)
UNAME_ARCH=$(shell uname -m)

# By default, archy is backwards compatible with uname for flags `-s` and `-m`
# So the following can be `uname -s` or `uname -m` respectively and this should still work
OS?=$(shell bin/archy -s)
ARCH?=$(shell bin/archy -m)

GOOS?=$(shell bin/archy -s --goruntime)
GOARCH?=$(shell bin/archy -m --goruntime)

ARCHY_VERSION?=0.1.0
KO_VERSION?=0.10.0
YTT_VERSION?=0.40.1

clean:
	rm -rf bin/*

bin/archy-devel:
	go build -o bin/archy ./cmd/archy

bin/archy:
	curl -L https://github.com/xargs-dev/archy/releases/download/v${ARCHY_VERSION}/archy_${ARCHY_VERSION}_${UNAME_OS}_${UNAME_ARCH}.tar.gz | tar xzf - archy
	chmod +x archy
	mv archy bin/

bin/ko:
	curl -L https://github.com/google/ko/releases/download/v${KO_VERSION}/ko_${KO_VERSION}_${OS}_${ARCH}.tar.gz | tar xzf - ko
	chmod +x ./ko
	mv ko bin/

bin/ytt:
	curl -L -o bin/ytt https://github.com/vmware-tanzu/carvel-ytt/releases/download/v${YTT_VERSION}/ytt-${GOOS}-${GOARCH}
	chmod +x bin/ytt

koversion: bin/ko
koversion:
	bin/ko version
