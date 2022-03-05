# `archy`

`archy` is an simple binary to determine current kernel and machine architecture, which wraps `uname` and alternatively can read from [Go `runtime` stdlib](https://pkg.go.dev/runtime) for `GOOS` and `GOARCH`.

## But, why!?

Some project might use Makefile to download their dependencies of GitHub releases. They often look like such:

```
https://github.com/google/ko/releases/download/v0.10.0/ko_0.10.0_Linux_x86_64.tar.gz
https://github.com/vmware-tanzu/carvel-ytt/releases/download/v0.40.1/ytt-darwin-amd64
```

In said Makefiles, usually the OS and architecture is variable-escaped, to match developer workstation or CI machine.
Unfortunately, that may not be as easy for the second link above, as the `uname` output is as such:

```
❯ uname -s
Darwin # URLs are case-sensitive

❯ uname -m
x86_64 # Needs amd64
```

One solution is to rely on `go env {GOOS,GOARCH}`, which outputs `darwin` and `amd64` respectively, but it requires Go being installed, which is not applicable in non-Go projects.

`archy` removes that dependency on Go binary being present – you can find a working example in [`Makefile`](Makefile).
