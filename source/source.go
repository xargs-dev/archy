package source

import (
	"context"

	"go.xargs.dev/archy"
	"go.xargs.dev/archy/option"
)

type Source interface {
	Values(context.Context, *option.Option, *archy.Values) error
}

func Use(str string) Source {
	switch str {
	case "go-runtime", "runtime", "goruntime":
		return &GoRuntime{}
	default:
		return &Uname{}
	}
}
