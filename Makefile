# By default, archy is backwards compatible with uname for flags `-s` and `-m`
# So the following can be `uname -s` or `uname -m` respectively and this should still work
OS?=$(shell bin/archy -s)
ARCH?=$(shell bin/archy -m)

GOOS?=$(shell bin/archy -s --goruntime)
GOARCH?=$(shell bin/archy -m --goruntime)

ARCHY_VERSION?=0.0.1
KO_VERSION?=0.10.0
YTT_VERSION?=0.40.1

bin/archy-devel:
	go build -o bin/archy ./cmd/archy

bin/archy:
	curl -Lo bin/archy https://github.com/xargs-dev/archy/release/download/v${ARCHY_VERSION}/archy-${OS}-${ARCH}
	chmod +x bin/archy

bin/ko:
	curl -L https://github.com/google/ko/releases/download/v${KO_VERSION}/ko_${KO_VERSION}_${OS}_${ARCH}.tar.gz | tar xzf - ko
	chmod +x ./ko
	mv ko bin/ko

bin/ytt:
	curl -L -o bin/ytt https://github.com/vmware-tanzu/carvel-ytt/releases/download/v${YTT_VERSION}/ytt-${GOOS}-${GOARCH}
	chmod +x bin/ytt

koversion: bin/ko
koversion:
	bin/ko version
