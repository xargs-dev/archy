package source

import (
	"context"
	"runtime"

	"go.xargs.dev/archy"
	"go.xargs.dev/archy/option"
)

type GoRuntime struct{}

func (*GoRuntime) Values(_ context.Context, o *option.Option, v *archy.Values) error {
	v.Source = "goruntime"
	if o.KernelName {
		v.KernelName = runtime.GOOS
	}
	if o.Machine {
		v.Machine = runtime.GOARCH
	}
	return nil
}
