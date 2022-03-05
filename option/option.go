package option

import (
	"github.com/spf13/pflag"
	"go.xargs.dev/archy/uname"
)

type Option struct {
	KernelName bool
	Machine    bool

	Source string

	JSON bool
}

func ParseFlags() *Option {
	o := &Option{}

	// Ensure future additions (if any) preserves backwards compatibility with "uname"
	pflag.BoolVarP(
		&o.KernelName,
		uname.FlagKernelName,
		"s",
		false,
		"print the kernel name (e.g. Linux, Darwin)",
	)
	pflag.BoolVarP(
		&o.Machine,
		uname.FlagMachine,
		"m",
		false,
		"print the machine hardware name (e.g. x86_64, arm64)",
	)
	pflag.StringVar(
		&o.Source,
		"source",
		"uname",
		"data source, options: uname, goruntime",
	)
	pflag.BoolVar(
		&o.JSON,
		"json",
		false,
		"print output as JSON",
	)

	goruntime := pflag.Bool("goruntime", false, "Use values from goruntime (i.e. GOOS, GOARCH), overrides --source")

	pflag.Parse()

	if *goruntime {
		o.Source = "goruntime"
	}

	if o.none() {
		// Per uname (GNU coreutils) 9.0:
		//   With no OPTION, same as -s.
		o.KernelName = true
	}

	return o
}

func (o *Option) none() bool {
	switch {
	case o.KernelName,
		o.Machine:
		return false
	default:
		return true
	}
}
